import { useApolloClient } from '@apollo/client'
import { useAuthContext } from '@authentication/AuthContext'
import { Avatar } from '@components/Avatar/Avatar'
import { Button } from '@components/Button'
import JsonViewer from '@components/JsonViewer/JsonViewer'
import { KeyboardShortcut } from '@components/KeyboardShortcut/KeyboardShortcut'
import { LinkButton } from '@components/LinkButton'
import LoadingBox from '@components/LoadingBox'
import { Skeleton } from '@components/Skeleton/Skeleton'
import {
	GetErrorInstanceDocument,
	useGetErrorInstanceQuery,
} from '@graph/hooks'
import { GetErrorGroupQuery, GetErrorObjectQuery } from '@graph/operations'
import {
	ErrorInstance as ErrorInstanceType,
	ErrorObject,
	Maybe,
	ReservedLogKey,
} from '@graph/schemas'
import {
	Box,
	Callout,
	Heading,
	IconSolidExternalLink,
	IconSolidLogs,
	IconSolidPlay,
	Text,
	Tooltip,
} from '@highlight-run/ui'
import { useProjectId } from '@hooks/useProjectId'
import ErrorStackTrace from '@pages/ErrorsV2/ErrorStackTrace/ErrorStackTrace'
import {
	DEFAULT_LOGS_OPERATOR,
	LogsSearchParam,
	stringifyLogsQuery,
} from '@pages/LogsPage/SearchForm/utils'
import { PlayerSearchParameters } from '@pages/Player/PlayerHook/utils'
import {
	getDisplayNameAndField,
	getIdentifiedUserProfileImage,
	getUserProperties,
} from '@pages/Sessions/SessionsFeedV3/MinimalSessionCard/utils/utils'
import analytics from '@util/analytics'
import { loadSession } from '@util/preload'
import { useParams } from '@util/react-router/useParams'
import { copyToClipboard } from '@util/string'
import { buildQueryURLString } from '@util/url/params'
import moment from 'moment'
import React, { useEffect, useState } from 'react'
import { useHotkeys } from 'react-hotkeys-hook'
import { createSearchParams, useNavigate } from 'react-router-dom'

const MAX_USER_PROPERTIES = 4
type Props = React.PropsWithChildren & {
	errorGroup: GetErrorGroupQuery['error_group']
}

const METADATA_LABELS: { [key: string]: string } = {
	os: 'OS',
	url: 'URL',
	id: 'ID',
} as const

const getLogsLink = (errorObject: ErrorObject): string => {
	const queryParams: LogsSearchParam[] = []
	let offsetStart = 1
	if (errorObject.session?.secure_id) {
		queryParams.push({
			key: ReservedLogKey.SecureSessionId,
			operator: DEFAULT_LOGS_OPERATOR,
			value: errorObject.session?.secure_id,
			offsetStart: offsetStart++,
		})
	}
	if (errorObject.trace_id) {
		queryParams.push({
			key: ReservedLogKey.TraceId,
			operator: DEFAULT_LOGS_OPERATOR,
			value: errorObject.trace_id,
			offsetStart: offsetStart++,
		})
	}
	const query = stringifyLogsQuery(queryParams)
	const logCursor = errorObject.log_cursor
	const params = createSearchParams({
		query,
		start_date: moment(errorObject.timestamp)
			.add(-5, 'minutes')
			.toISOString(),
		end_date: moment(errorObject.timestamp).add(5, 'minutes').toISOString(),
	})
	if (logCursor) {
		return `/${errorObject.project_id}/logs/${logCursor}?${params}`
	} else {
		return `/${errorObject.project_id}/logs?${params}`
	}
}

const getSessionLink = (errorObject: ErrorObject): string => {
	const params = createSearchParams({
		tsAbs: errorObject.timestamp,
		[PlayerSearchParameters.errorId]: errorObject.id,
	})
	return `/${errorObject.project_id}/sessions/${errorObject.session?.secure_id}?${params}`
}

const ErrorInstance: React.FC<Props> = ({ errorGroup }) => {
	const { error_object_id, error_secure_id } = useParams<{
		error_secure_id: string
		error_object_id: string
	}>()
	const { projectId } = useProjectId()
	const navigate = useNavigate()
	const client = useApolloClient()
	const { isLoggedIn } = useAuthContext()

	const { loading, data } = useGetErrorInstanceQuery({
		variables: {
			error_group_secure_id: String(errorGroup?.secure_id),
			error_object_id,
		},
		onCompleted: (data) => {
			const previousErrorObjectId = data?.error_instance?.previous_id
			const nextErrorObjectId = data?.error_instance?.next_id

			// Prefetch the next/previous error objects so they are in the cache.
			// Using client directly because the lazy query had issues with canceling
			// multiple requests: https://github.com/apollographql/apollo-client/issues/9755
			if (previousErrorObjectId) {
				client.query({
					query: GetErrorInstanceDocument,
					variables: {
						error_group_secure_id: String(errorGroup?.secure_id),
						error_object_id: previousErrorObjectId,
					},
				})
			}

			if (nextErrorObjectId) {
				client.query({
					query: GetErrorInstanceDocument,
					variables: {
						error_group_secure_id: String(errorGroup?.secure_id),
						error_object_id: nextErrorObjectId,
					},
				})
			}

			// Prefetch session data.
			if (data?.error_instance?.error_object?.session) {
				loadSession(data.error_instance.error_object.session.secure_id)
			}
		},
	})

	const errorInstance = data?.error_instance

	useEffect(() => analytics.page(), [])

	useHotkeys(']', () => goToErrorInstance(errorInstance?.next_id, 'next'), [
		errorInstance?.next_id,
	])
	useHotkeys(
		'[',
		() => goToErrorInstance(errorInstance?.previous_id, 'previous'),
		[errorInstance?.previous_id],
	)

	const goToErrorInstance = (
		errorInstanceId: Maybe<string> | undefined,
		direction: 'next' | 'previous',
	) => {
		if (!errorInstanceId || Number(errorInstanceId) === 0) {
			return
		}

		navigate({
			pathname: `/${projectId}/errors/${error_secure_id}/instances/${errorInstanceId}`,
			search: window.location.search,
		})

		analytics.track('Viewed error instance', {
			direction,
		})
	}

	if (!errorInstance || !errorInstance?.error_object) {
		if (!loading) return null

		return (
			<Box id="error-instance-container">
				<Box my="28" display="flex" justifyContent="space-between">
					<Box display="flex" flexDirection="column" gap="16">
						<Heading level="h4">Error Instance</Heading>
					</Box>

					<Box>
						<Box display="flex" gap="8" alignItems="center">
							<Button
								disabled={true}
								kind="secondary"
								emphasis="low"
								trackingId="errorInstanceOlder"
							>
								Older
							</Button>
							<Box
								borderRight="secondary"
								style={{ height: 18 }}
							/>
							<Button
								disabled={true}
								kind="secondary"
								emphasis="low"
								trackingId="errorInstanceNewer"
							>
								Newer
							</Button>
							<Button
								kind="primary"
								emphasis="high"
								disabled={true}
								iconLeft={<IconSolidLogs />}
								trackingId="errorInstanceShowLogs"
							>
								Show logs
							</Button>
							<Button
								kind="primary"
								emphasis="high"
								disabled={true}
								iconLeft={<IconSolidPlay />}
								trackingId="errorInstanceShowSession"
							>
								Show session
							</Button>
						</Box>
					</Box>
				</Box>

				<Box
					display="flex"
					flexDirection={{ desktop: 'row', mobile: 'column' }}
					mb="40"
					gap="40"
				>
					<div style={{ flexBasis: 0, flexGrow: 1 }}>
						<Box>
							<Box bb="secondary" pb="20" my="12">
								<Text weight="bold" size="large">
									Instance metadata
								</Text>
							</Box>
							<LoadingBox height={128} />
						</Box>
					</div>

					<div style={{ flexBasis: 0, flexGrow: 1 }}>
						<Box width="full">
							<Box pb="20" mt="12">
								<Text weight="bold" size="large">
									User details
								</Text>
							</Box>
							<LoadingBox height={128} />
						</Box>
					</div>
				</Box>

				<Text size="large" weight="bold">
					Stack trace
				</Text>
				<Box bt="secondary" mt="12" pt="16">
					<Skeleton count={10} />
				</Box>
			</Box>
		)
	}

	const errorObject =
		errorInstance.error_object as ErrorInstanceType['error_object']

	return (
		<Box id="error-instance-container">
			<Box my="28" display="flex" justifyContent="space-between">
				<Box display="flex" flexDirection="column" gap="16">
					<Heading level="h4">Error Instance</Heading>
				</Box>

				<Box>
					<Box display="flex" gap="8" alignItems="center">
						<Tooltip
							trigger={
								<Button
									onClick={() => {
										goToErrorInstance(
											errorInstance.previous_id,
											'previous',
										)
									}}
									disabled={
										Number(errorInstance.previous_id) === 0
									}
									kind="secondary"
									emphasis="low"
									trackingId="errorInstanceOlder"
								>
									Older
								</Button>
							}
						>
							<KeyboardShortcut label="Previous" shortcut="[" />
						</Tooltip>

						<Box borderRight="secondary" style={{ height: 18 }} />
						<Tooltip
							trigger={
								<Button
									onClick={() => {
										goToErrorInstance(
											errorInstance.next_id,
											'next',
										)
									}}
									disabled={
										Number(errorInstance.next_id) === 0
									}
									kind="secondary"
									emphasis="low"
									trackingId="errorInstanceNewer"
								>
									Newer
								</Button>
							}
						>
							<KeyboardShortcut label="Next" shortcut="]" />
						</Tooltip>
						<LinkButton
							kind="primary"
							emphasis="high"
							to={getLogsLink(errorObject)}
							disabled={!isLoggedIn || !errorObject.session}
							trackingId="error-related_logs_link"
						>
							<Box
								display="flex"
								alignItems="center"
								flexDirection="row"
								gap="4"
							>
								<IconSolidLogs />
								Show logs
							</Box>
						</LinkButton>
						<LinkButton
							kind="primary"
							emphasis="high"
							to={getSessionLink(errorObject)}
							disabled={!isLoggedIn || !errorObject.session}
							trackingId="error-related_session_link"
						>
							<Box
								display="flex"
								alignItems="center"
								flexDirection="row"
								gap="4"
							>
								<IconSolidPlay />
								Show session
							</Box>
						</LinkButton>
					</Box>
				</Box>
			</Box>

			<Box
				display="flex"
				flexDirection={{ desktop: 'row', mobile: 'column' }}
				mb="40"
				gap="40"
			>
				<div style={{ flexBasis: 0, flexGrow: 1 }}>
					<Metadata errorObject={errorObject} />
				</div>

				<div style={{ flexBasis: 0, flexGrow: 1 }}>
					<User errorObject={errorObject} />
				</div>
			</Box>

			{errorGroup?.type === 'console.error' &&
				errorGroup.event.length > 1 && (
					<>
						<Text size="large" weight="bold">
							Error event data
						</Text>

						<Box bt="secondary" my="12" py="16">
							<JsonViewer src={errorGroup.event} collapsed={1} />
						</Box>
					</>
				)}

			{(errorObject.stack_trace !== '' &&
				errorObject.stack_trace !== 'null') ||
			errorObject.structured_stack_trace?.length ? (
				<>
					<Text size="large" weight="bold">
						Stack trace
					</Text>
					<Box bt="secondary" mt="12" pt="16">
						<ErrorStackTrace
							errorObject={
								errorObject as ErrorInstanceType['error_object']
							}
						/>
					</Box>
				</>
			) : null}
		</Box>
	)
}

const Metadata: React.FC<{
	errorObject?: GetErrorObjectQuery['error_object']
}> = ({ errorObject }) => {
	if (!errorObject) {
		return null
	}

	let customProperties: any
	try {
		if (errorObject.payload) {
			customProperties = JSON.parse(errorObject.payload)
		}
	} catch (e) {}

	const metadata = [
		{ key: 'environment', label: errorObject?.environment },
		{ key: 'browser', label: errorObject?.browser },
		{ key: 'os', label: errorObject?.os },
		{ key: 'url', label: errorObject?.url },
		{ key: 'timestamp', label: errorObject?.timestamp },
		{
			key: 'Custom Properties',
			label: customProperties ? (
				<JsonViewer
					collapsed={true}
					src={customProperties}
					name="Custom Properties"
				/>
			) : undefined,
		},
	].filter((t) => Boolean(t.label))

	return (
		<Box>
			<Box bb="secondary" pb="20" my="12">
				<Text weight="bold" size="large">
					Instance metadata
				</Text>
			</Box>

			<Box>
				{metadata.map((meta) => {
					const value =
						meta.key === 'timestamp'
							? moment(meta.label as string).format(
									'M/D/YY h:mm:ss.SSS A',
							  )
							: meta.label
					return (
						<Box display="flex" gap="6" key={meta.key}>
							<Box
								py="10"
								cursor="pointer"
								onClick={() => copyToClipboard(meta.key)}
								style={{ width: '33%' }}
							>
								<Text
									color="n11"
									transform="capitalize"
									align="left"
									lines="1"
								>
									{METADATA_LABELS[meta.key] ??
										meta.key.replace('_', ' ')}
								</Text>
							</Box>
							<Box
								cursor="pointer"
								py="10"
								onClick={() => {
									if (typeof value === 'string') {
										value && copyToClipboard(value)
									}
								}}
								style={{ width: '67%' }}
							>
								<Text
									align="left"
									break="word"
									lines={
										typeof value === 'string'
											? '4'
											: undefined
									}
									title={String(value)}
								>
									{value}
								</Text>
							</Box>
						</Box>
					)
				})}
			</Box>
		</Box>
	)
}

const User: React.FC<{
	errorObject?: GetErrorObjectQuery['error_object']
}> = ({ errorObject }) => {
	const navigate = useNavigate()
	const { projectId } = useProjectId()
	const { isLoggedIn } = useAuthContext()
	const [truncated, setTruncated] = useState(true)

	if (!errorObject?.session) {
		return (
			<Box width="full">
				<Box pb="20" mt="12">
					<Text weight="bold" size="large">
						User details
					</Text>
				</Box>
				<Callout title="We didn't find a session for this error">
					<Box>
						<Text size="small" weight="medium" color="moderate">
							We weren't able to match this error to a session.
							This error was either thrown in isolation or you
							aren't mapping errors to sessions.
						</Text>
					</Box>
					<Box display="flex">
						<LinkButton
							kind="secondary"
							to={`/${projectId}/setup/backend`}
							trackingId="error-mapping-setup"
							target="_blank"
						>
							Backend SDK setup
						</LinkButton>
						<LinkButton
							kind="secondary"
							to="https://www.highlight.io/docs/getting-started/frontend-backend-mapping"
							trackingId="error-mapping-docs"
							emphasis="low"
							target="_blank"
						>
							Learn more
						</LinkButton>
					</Box>
				</Callout>
			</Box>
		)
	}

	const { session } = errorObject
	const userProperties = getUserProperties(session?.user_properties)
	const [displayName, field] = getDisplayNameAndField(session)
	const avatarImage = getIdentifiedUserProfileImage(session)
	const userDisplayPropertyKeys = Object.keys(userProperties)
		.filter((k) => k !== 'avatar')
		.slice(
			0,
			truncated
				? MAX_USER_PROPERTIES - 1
				: Object.keys(userProperties).length,
		)

	const truncateable =
		Object.keys(userProperties).length > MAX_USER_PROPERTIES
	const location = [session?.city, session?.state, session?.country]
		.filter(Boolean)
		.join(', ')

	return (
		<Box width="full">
			<Box pb="20" mt="12">
				<Text weight="bold" size="large">
					User details
				</Text>
			</Box>
			<Box border="secondary" borderRadius="6">
				<Box
					bb="secondary"
					py="8"
					px="12"
					alignItems="center"
					display="flex"
					justifyContent="space-between"
					gap="4"
				>
					<Box alignItems="center" display="flex" gap="8">
						<Avatar
							seed={displayName}
							style={{ height: 28, width: 28 }}
							customImage={avatarImage}
						/>
						<Text lines="1">{displayName}</Text>
					</Box>

					<Box flexShrink={0} display="flex">
						<Button
							kind="secondary"
							emphasis="high"
							iconRight={<IconSolidExternalLink />}
							disabled={!isLoggedIn}
							onClick={() => {
								if (!isLoggedIn) {
									return
								}

								const searchParams: any = {}
								if (session.identifier && field !== null) {
									searchParams[`user_${field}`] = displayName
								} else if (session?.fingerprint) {
									searchParams.device_id = String(
										session.fingerprint,
									)
								}

								navigate({
									pathname: `/${projectId}/sessions`,
									search: buildQueryURLString(searchParams),
								})
							}}
							trackingId="errorInstanceAllSessionsForuser"
						>
							All sessions for this user
						</Button>
					</Box>
				</Box>

				<Box py="8" px="12">
					<Box display="flex" flexDirection="column">
						<Box gap="16">
							{userDisplayPropertyKeys.map((key) => (
								<Box display="flex" gap="6" key={key}>
									<Box
										py="10"
										overflow="hidden"
										onClick={() => copyToClipboard(key)}
										style={{ width: '33%' }}
									>
										<Text
											color="n11"
											align="left"
											transform="capitalize"
											lines="1"
											title={key}
										>
											{METADATA_LABELS[key] ?? key}
										</Text>
									</Box>

									<Box
										py="10"
										display="flex"
										overflow="hidden"
										onClick={() =>
											copyToClipboard(userProperties[key])
										}
										title={userProperties[key]}
										style={{ width: '67%' }}
									>
										<Text lines="1" as="span">
											{userProperties[key]}
										</Text>
									</Box>
								</Box>
							))}

							<Box display="flex" alignItems="center" gap="6">
								<Box py="10" style={{ width: '33%' }}>
									<Text color="n11" align="left">
										Location
									</Text>
								</Box>

								<Box py="10" style={{ width: '67%' }}>
									<Text>{location}</Text>
								</Box>
							</Box>
						</Box>
						{truncateable && (
							<Box>
								<Button
									onClick={() => setTruncated(!truncated)}
									kind="secondary"
									emphasis="medium"
									size="xSmall"
									trackingId="errorInstanceToggleProperties"
								>
									Show {truncated ? 'more' : 'less'}
								</Button>
							</Box>
						)}
					</Box>
				</Box>
			</Box>
		</Box>
	)
}

export default ErrorInstance
