package graph

import (
	"context"
	"encoding/json"
	"github.com/highlight-run/highlight/backend/timeseries"
	"github.com/highlight-run/workerpool"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"testing"
	"time"

	e "github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/highlight-run/highlight/backend/model"
	publicModelInput "github.com/highlight-run/highlight/backend/public-graph/graph/model"
	"github.com/highlight-run/highlight/backend/util"
)

var DB *gorm.DB

// Gets run once; M.run() calls the tests in this file.
func TestMain(m *testing.M) {
	dbName := "highlight_testing_db"
	testLogger := log.WithContext(context.TODO()).WithFields(log.Fields{"DB_HOST": os.Getenv("PSQL_HOST"), "DB_NAME": dbName})
	var err error
	DB, err = util.CreateAndMigrateTestDB("highlight_testing_db")
	if err != nil {
		testLogger.Error(e.Wrap(err, "error creating testdb"))
	}
	code := m.Run()
	os.Exit(code)
}

func TestProcessBackendPayloadImpl(t *testing.T) {
	trpcTraceStr := "[{\"columnNumber\":11,\"lineNumber\":80,\"fileName\":\"/workspace/src/trpc/instance.ts\",\"source\":\"    at /workspace/src/trpc/instance.ts:80:11\",\"lineContent\":\"    throw new TRPCError({\\n\",\"linesBefore\":\"        organizationId,\\n        supabaseAccessToken,\\n      },\\n    });\\n  } catch (error) {\\n\",\"linesAfter\":\"      code: \\\"UNAUTHORIZED\\\",\\n    });\\n  }\\n});\\n\\n\"},{\"columnNumber\":38,\"lineNumber\":421,\"fileName\":\"/workspace/node_modules/@trpc/server/dist/index.js\",\"functionName\":\"callRecursive\",\"source\":\"    at callRecursive (/workspace/node_modules/@trpc/server/dist/index.js:421:38)\",\"lineContent\":\"                const result = await middleware({\\n\",\"linesBefore\":\"            ctx: opts.ctx\\n        })=\u003e{\\n            try {\\n                // eslint-disable-next-line @typescript-eslint/no-non-null-assertion\\n                const middleware = _def.middlewares[callOpts.index];\\n\",\"linesAfter\":\"                    ctx: callOpts.ctx,\\n                    type: opts.type,\\n                    path: opts.path,\\n                    rawInput: opts.rawInput,\\n                    meta: _def.meta,\\n\"},{\"columnNumber\":30,\"lineNumber\":449,\"fileName\":\"/workspace/node_modules/@trpc/server/dist/index.js\",\"functionName\":\"resolve\",\"source\":\"    at resolve (/workspace/node_modules/@trpc/server/dist/index.js:449:30)\",\"lineContent\":\"        const result = await callRecursive();\\n\",\"linesBefore\":\"                    marker: middlewareMarker\\n                };\\n            }\\n        };\\n        // there's always at least one \\\"next\\\" since we wrap this.resolver in a middleware\\n\",\"linesAfter\":\"        if (!result) {\\n            throw new TRPCError.TRPCError({\\n                code: 'INTERNAL_SERVER_ERROR',\\n                message: 'No result from middlewares - did you forget to `return next()`?'\\n            });\\n\"},{\"columnNumber\":12,\"lineNumber\":228,\"fileName\":\"/workspace/node_modules/@trpc/server/dist/config-7b65d7da.js\",\"functionName\":\"Object.callProcedure\",\"source\":\"    at Object.callProcedure (/workspace/node_modules/@trpc/server/dist/config-7b65d7da.js:228:12)\",\"lineContent\":\"    return procedure(opts);\\n\",\"linesBefore\":\"            code: 'NOT_FOUND',\\n            message: `No \\\"${type}\\\"-procedure on path \\\"${path}\\\"`\\n        });\\n    }\\n    const procedure = opts.procedures[path];\\n\",\"linesAfter\":\"}\\n\\n/**\\n * The default check to see if we're in a server\\n */ const isServerDefault = typeof window === 'undefined' || 'Deno' in window || globalThis.process?.env?.NODE_ENV === 'test' || !!globalThis.process?.env?.JEST_WORKER_ID;\\n\"},{\"columnNumber\":45,\"lineNumber\":125,\"fileName\":\"/workspace/node_modules/@trpc/server/dist/resolveHTTPResponse-83d9b5ff.js\",\"source\":\"    at /workspace/node_modules/@trpc/server/dist/resolveHTTPResponse-83d9b5ff.js:125:45\",\"lineContent\":\"                const output = await config.callProcedure({\\n\",\"linesBefore\":\"        };\\n        const inputs = getInputs();\\n        const rawResults = await Promise.all(paths.map(async (path, index)=\u003e{\\n            const input = inputs[index];\\n            try {\\n\",\"linesAfter\":\"                    procedures: router._def.procedures,\\n                    path,\\n                    rawInput: input,\\n                    ctx,\\n                    type\\n\"},{\"columnNumber\":52,\"lineNumber\":122,\"fileName\":\"/workspace/node_modules/@trpc/server/dist/resolveHTTPResponse-83d9b5ff.js\",\"functionName\":\"Object.resolveHTTPResponse\",\"source\":\"    at Object.resolveHTTPResponse (/workspace/node_modules/@trpc/server/dist/resolveHTTPResponse-83d9b5ff.js:122:52)\",\"lineContent\":\"        const rawResults = await Promise.all(paths.map(async (path, index)=\u003e{\\n\",\"linesBefore\":\"                input[k] = value;\\n            }\\n            return input;\\n        };\\n        const inputs = getInputs();\\n\",\"linesAfter\":\"            const input = inputs[index];\\n            try {\\n                const output = await config.callProcedure({\\n                    procedures: router._def.procedures,\\n                    path,\\n\"},{\"columnNumber\":5,\"lineNumber\":96,\"fileName\":\"node:internal/process/task_queues\",\"functionName\":\"processTicksAndRejections\",\"source\":\"    at processTicksAndRejections (node:internal/process/task_queues:96:5)\"},{\"columnNumber\":20,\"lineNumber\":53,\"fileName\":\"/workspace/node_modules/@trpc/server/dist/nodeHTTPRequestHandler-e6a535cb.js\",\"functionName\":\"Object.nodeHTTPRequestHandler\",\"source\":\"    at Object.nodeHTTPRequestHandler (/workspace/node_modules/@trpc/server/dist/nodeHTTPRequestHandler-e6a535cb.js:53:20)\",\"lineContent\":\"    const result = await resolveHTTPResponse.resolveHTTPResponse({\\n\",\"linesBefore\":\"        method: opts.req.method,\\n        headers: opts.req.headers,\\n        query,\\n        body: bodyResult.ok ? bodyResult.data : undefined\\n    };\\n\",\"linesAfter\":\"        batching: opts.batching,\\n        responseMeta: opts.responseMeta,\\n        path,\\n        createContext,\\n        router,\\n\"}]"
	util.RunTestWithDBWipe(t, "trpc test", DB, func(t *testing.T) {
		r := &Resolver{AlertWorkerPool: workerpool.New(1), DB: DB, TDB: timeseries.New(context.TODO())}
		r.ProcessBackendPayloadImpl(context.Background(), nil, nil, []*publicModelInput.BackendErrorObjectInput{{
			SessionSecureID: nil,
			RequestID:       nil,
			TraceID:         nil,
			SpanID:          nil,
			Event:           "dummy event",
			Type:            "",
			URL:             "",
			Source:          "",
			StackTrace:      trpcTraceStr,
			Timestamp:       time.Time{},
			Payload:         nil,
		}})
		var result *model.ErrorObject
		r.DB.Model(&model.ErrorObject{}).Where(&model.ErrorObject{Event: "dummy event"}).First(&result)
		if *result.StackTrace != trpcTraceStr {
			t.Fatal("stacktrace changed after processing")
		}
	})
}

func TestHandleErrorAndGroup(t *testing.T) {
	// construct table of sub-tests to run
	longTraceStr := `[{"functionName":"is","args":null,"fileName":null,"lineNumber":null,"columnNumber":null,"isEval":null,"isNative":null,"source":null},{"functionName":"longer","args":null,"fileName":null,"lineNumber":null,"columnNumber":null,"isEval":null,"isNative":null,"source":null},{"functionName":"trace","args":null,"fileName":null,"lineNumber":null,"columnNumber":null,"isEval":null,"isNative":null,"source":null}]`
	shortTraceStr := `[{"functionName":"a","args":null,"fileName":null,"lineNumber":null,"columnNumber":null,"isEval":null,"isNative":null,"source":null},{"functionName":"short","args":null,"fileName":null,"lineNumber":null,"columnNumber":null,"isEval":null,"isNative":null,"source":null}]`
	tests := map[string]struct {
		errorsToInsert      []model.ErrorObject
		expectedErrorGroups []model.ErrorGroup
	}{
		"test two errors with same environment but different case": {
			errorsToInsert: []model.ErrorObject{
				{
					Event:       "error",
					ProjectID:   1,
					Environment: "dev",
					Model:       model.Model{CreatedAt: time.Date(2000, 8, 1, 0, 0, 0, 0, time.UTC), ID: 1},
				},
				{
					Event:       "error",
					ProjectID:   1,
					Environment: "dEv",
					Model:       model.Model{CreatedAt: time.Date(2000, 8, 1, 0, 0, 0, 0, time.UTC), ID: 2},
				},
			},
			expectedErrorGroups: []model.ErrorGroup{
				{
					Event:        "error",
					ProjectID:    1,
					State:        model.ErrorGroupStates.OPEN,
					Environments: `{"dev":2}`,
				},
			},
		},
		"test two errors with different environment": {
			errorsToInsert: []model.ErrorObject{
				{
					Event:       "error",
					ProjectID:   1,
					Environment: "dev",
					Model:       model.Model{CreatedAt: time.Date(2000, 8, 1, 0, 0, 0, 0, time.UTC), ID: 1},
				},
				{
					Event:       "error",
					ProjectID:   1,
					Environment: "prod",
					Model:       model.Model{CreatedAt: time.Date(2000, 8, 1, 0, 0, 0, 0, time.UTC), ID: 2},
				},
			},
			expectedErrorGroups: []model.ErrorGroup{
				{
					Event:        "error",
					ProjectID:    1,
					State:        model.ErrorGroupStates.OPEN,
					Environments: `{"dev":1,"prod":1}`,
				},
			},
		},
		"two errors, one with empty environment": {
			errorsToInsert: []model.ErrorObject{
				{
					ProjectID:   1,
					Environment: "dev",
					Model:       model.Model{CreatedAt: time.Date(2000, 8, 1, 0, 0, 0, 0, time.UTC), ID: 1},
					Event:       "error",
				},
				{
					Event:     "error",
					ProjectID: 1,
					Model:     model.Model{CreatedAt: time.Date(2000, 8, 1, 0, 0, 0, 0, time.UTC), ID: 2},
				},
			},
			expectedErrorGroups: []model.ErrorGroup{
				{
					Event:        "error",
					ProjectID:    1,
					State:        model.ErrorGroupStates.OPEN,
					Environments: `{"dev":1}`,
				},
			},
		},
		"test longer error stack first": {
			errorsToInsert: []model.ErrorObject{
				{
					Event:      "error",
					ProjectID:  1,
					Model:      model.Model{CreatedAt: time.Date(2000, 8, 1, 0, 0, 0, 0, time.UTC), ID: 1},
					StackTrace: &longTraceStr,
				},
				{
					Event:      "error",
					ProjectID:  1,
					Model:      model.Model{CreatedAt: time.Date(2000, 8, 1, 0, 0, 0, 0, time.UTC), ID: 2},
					StackTrace: &shortTraceStr,
				},
			},
			expectedErrorGroups: []model.ErrorGroup{
				{
					Event:            "error",
					ProjectID:        1,
					StackTrace:       shortTraceStr,
					State:            model.ErrorGroupStates.OPEN,
					Environments:     `{}`,
					MappedStackTrace: util.MakeStringPointer("null"),
				},
			},
		},
		"test shorter error stack first": {
			errorsToInsert: []model.ErrorObject{
				{
					Event:      "error",
					ProjectID:  1,
					Model:      model.Model{CreatedAt: time.Date(2000, 8, 1, 0, 0, 0, 0, time.UTC), ID: 1},
					StackTrace: &shortTraceStr,
				},
				{
					Event:      "error",
					ProjectID:  1,
					Model:      model.Model{CreatedAt: time.Date(2000, 8, 1, 0, 0, 0, 0, time.UTC), ID: 2},
					StackTrace: &longTraceStr,
				},
			},
			expectedErrorGroups: []model.ErrorGroup{
				{
					Event:            "error",
					ProjectID:        1,
					StackTrace:       longTraceStr,
					Environments:     `{}`,
					State:            model.ErrorGroupStates.OPEN,
					MappedStackTrace: util.MakeStringPointer("null"),
				},
			},
		},
	}
	// run tests
	for name, tc := range tests {
		util.RunTestWithDBWipe(t, name, DB, func(t *testing.T) {
			r := &Resolver{DB: DB}
			receivedErrorGroups := make(map[string]model.ErrorGroup)
			for _, errorObj := range tc.errorsToInsert {
				var frames []*publicModelInput.StackFrameInput
				if errorObj.StackTrace != nil {
					if err := json.Unmarshal([]byte(*errorObj.StackTrace), &frames); err != nil {
						t.Fatal(e.Wrap(err, "error unmarshalling error stack trace frames"))
					}
				}
				errorGroup, err := r.HandleErrorAndGroup(context.TODO(), &errorObj, "", frames, nil, 1)
				if err != nil {
					t.Fatal(e.Wrap(err, "error handling error and group"))
				}
				if errorGroup != nil {
					id := strconv.Itoa(errorGroup.ID)
					receivedErrorGroups[id] = *errorGroup
				}
			}
			var i int
			for _, errorGroup := range receivedErrorGroups {
				isEqual, diff, err := model.AreModelsWeaklyEqual(&errorGroup, &tc.expectedErrorGroups[i])
				if err != nil {
					t.Fatal(e.Wrap(err, "error comparing two error groups"))
				}
				if !isEqual {
					t.Fatalf("received error group not equal to expected error group. diff: %+v", diff)
				}
				i++
			}
		})
	}
}

func TestResolver_isExcludedError(t *testing.T) {
	r := &Resolver{}
	assert.False(t, r.isExcludedError(context.Background(), []string{}, "", 1))
	assert.True(t, r.isExcludedError(context.Background(), []string{}, "[{}]", 2))
	assert.True(t, r.isExcludedError(context.Background(), []string{".*a+.*"}, "foo bar baz", 3))
	assert.False(t, r.isExcludedError(context.Background(), []string{"("}, "foo bar baz", 4))
}
