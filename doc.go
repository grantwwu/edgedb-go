// This source file is part of the EdgeDB open source project.
//
// Copyright 2020-present EdgeDB Inc. and the EdgeDB authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package edgedb is the official Go EdgeDB driver. https://edgedb.com
//
// Typical usage looks like this:
//
//   import (
//       "context"
//       "log"
//
//       "github.com/edgedb/edgedb-go"
//   )
//
//   opts := edgedb.Options{
//       MinConns: 1,
//       MaxConns: 4,
//   }
//
//   func main() {
//       ctx := context.Background()
//       pool, err := edgedb.ConnectDSN(ctx, "my_instance", opts)
//       if err != nil {
//           log.Fatal(err)
//       }
//       defer pool.Close()
//
//       var (
//           age int64 = 21
//           users []struct{
//               ID edgedb.UUID `edgedb:"id"`
//               Name string    `edgedb:"name"`
//           }
//       )
//
//       query := "SELECT User{name} WHERE .age = <int64>$0"
//       err = pool.Query(ctx, query, &users, age)
//       ...
//   }
//
// You can also connect to a database using a DSN:
//
//   url := "edgedb://edgedb@localhost/edgedb"
//   pool, err := edgedb.ConnectDSN(ctx, url, opts)
//
// Or you can use Option fields.
//
//   opts := edgedb.Options{
//       Database: "edgedb",
//       User:     "edgedb",
//       MinConns: 1,
//       MaxConns: 4,
//   }
//
//   pool, err := edgedb.Connect(ctx, opts)
//
// Pooling
//
// Most use cases will benefit from the concurrency safe pool implementation
// returned from Connect() and ConnectDSN(). Pool.RetryingTx() or Pool.RawTx()
// will give you access to a single connection.
//
// Errors
//
// edgedb never returns underlying errors directly.
// If you are checking for things like context expiration
// use errors.Is() or errors.As().
//
//   err := pool.Query(...)
//   if errors.Is(err, context.Canceled) { ... }
//
// Most errors returned by the edgedb package will satisfy the edgedb.Error
// interface which has methods for introspecting.
//
//   err := pool.Query(...)
//
//   var edbErr edgedb.Error
//   if errors.As(err, &edbErr) && edbErr.Category(edgedb.NoDataError){
//       ...
//   }
//
// Datatypes
//
// The following list shows the marshal/unmarshal
// mapping between EdgeDB types and go types:
//
//   EdgeDB                   Go
//   ---------                ---------
//   Set                      []anytype
//   array<anytype>           []anytype
//   tuple                    struct
//   named tuple              struct
//   Object                   struct
//   bool                     bool, edgedb.OptionalBool
//   bytes                    []byte, edgedb.OptionalBytes
//   str                      string, edgedb.OptionalStr
//   anyenum                  string, edgedb.OptionalStr
//   datetime                 time.Time, edgedb.OptionalDateTime
//   cal::local_datetime      edgedb.LocalDateTime,
//                            edgedb.OptionalLocalDateTime
//   cal::local_date          edgedb.LocalDate, edgedb.OptionalLocalDate
//   cal::local_time          edgedb.LocalTime, edgedb.OptionalLocalTime
//   duration                 time.Duration, edgedb.OptionalDuration
//   cal::relative_duraation  edgedb.RelativeDuration,
//                            edgedb.OptionalRelativeDuration
//   float32                  float32, edgedb.OptionalFloat32
//   float64                  float64, edgedb.OptionalFloat64
//   int16                    int16, edgedb.OptionalFloat16
//   int32                    int32, edgedb.OptionalInt16
//   int64                    int64, edgedb.OptionalInt64
//   uuid                     edgedb.UUID, edgedb.OptionalUUID
//   json                     []byte, edgedb.OptionalBytes
//   bigint                   *big.Int, edgedb.OptionalBigInt
//
//   decimal                  user defined (see Custom Marshalers)
//
// Shape fileds that are not required must use optional types for receiving
// query results. The edgedb.Optional struct can be embedded to make structs
// optional.
//
// Custom Marshalers
//
// Interfaces for user defined marshaler/unmarshalers  are documented in the
// internal/marshal package.
package edgedb
