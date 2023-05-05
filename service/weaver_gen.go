package service

// Code generated by "weaver generate". DO NOT EDIT.
import (
	"context"
	"errors"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"reflect"
	"time"
)

func init() {
	codegen.Register(codegen.Registration{
		Name:        "github.com/mendoncas/weaver-study/service/Books",
		Iface:       reflect.TypeOf((*Books)(nil)).Elem(),
		New:         func() any { return &books{} },
		LocalStubFn: func(impl any, tracer trace.Tracer) any { return books_local_stub{impl: impl.(Books), tracer: tracer} },
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return books_client_stub{stub: stub, registerBookMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/mendoncas/weaver-study/service/Books", Method: "RegisterBook"}), findBookByTitleMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/mendoncas/weaver-study/service/Books", Method: "FindBookByTitle"})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return books_server_stub{impl: impl.(Books), addLoad: addLoad}
		},
	})
	codegen.Register(codegen.Registration{
		Name:  "github.com/mendoncas/weaver-study/service/Details",
		Iface: reflect.TypeOf((*Details)(nil)).Elem(),
		New:   func() any { return &details{} },
		LocalStubFn: func(impl any, tracer trace.Tracer) any {
			return details_local_stub{impl: impl.(Details), tracer: tracer}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return details_client_stub{stub: stub, addBookDetailsMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/mendoncas/weaver-study/service/Details", Method: "AddBookDetails"}), findDetailsByBookIdMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/mendoncas/weaver-study/service/Details", Method: "FindDetailsByBookId"})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return details_server_stub{impl: impl.(Details), addLoad: addLoad}
		},
	})
	codegen.Register(codegen.Registration{
		Name:  "github.com/mendoncas/weaver-study/service/Reverser",
		Iface: reflect.TypeOf((*Reverser)(nil)).Elem(),
		New:   func() any { return &reverser{} },
		LocalStubFn: func(impl any, tracer trace.Tracer) any {
			return reverser_local_stub{impl: impl.(Reverser), tracer: tracer}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return reverser_client_stub{stub: stub, reverseMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/mendoncas/weaver-study/service/Reverser", Method: "Reverse"})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return reverser_server_stub{impl: impl.(Reverser), addLoad: addLoad}
		},
	})
	codegen.Register(codegen.Registration{
		Name:  "github.com/mendoncas/weaver-study/service/Reviews",
		Iface: reflect.TypeOf((*Reviews)(nil)).Elem(),
		New:   func() any { return &reviews{} },
		LocalStubFn: func(impl any, tracer trace.Tracer) any {
			return reviews_local_stub{impl: impl.(Reviews), tracer: tracer}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return reviews_client_stub{stub: stub, getAllBookReviewsMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/mendoncas/weaver-study/service/Reviews", Method: "GetAllBookReviews"})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return reviews_server_stub{impl: impl.(Reviews), addLoad: addLoad}
		},
	})
}

// Local stub implementations.

type books_local_stub struct {
	impl   Books
	tracer trace.Tracer
}

func (s books_local_stub) RegisterBook(ctx context.Context, a0 Book) (err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "service.Books.RegisterBook", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.RegisterBook(ctx, a0)
}

func (s books_local_stub) FindBookByTitle(ctx context.Context, a0 string) (r0 []byte, err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "service.Books.FindBookByTitle", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.FindBookByTitle(ctx, a0)
}

type details_local_stub struct {
	impl   Details
	tracer trace.Tracer
}

func (s details_local_stub) AddBookDetails(ctx context.Context, a0 string) (err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "service.Details.AddBookDetails", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.AddBookDetails(ctx, a0)
}

func (s details_local_stub) FindDetailsByBookId(ctx context.Context, a0 string) (err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "service.Details.FindDetailsByBookId", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.FindDetailsByBookId(ctx, a0)
}

type reverser_local_stub struct {
	impl   Reverser
	tracer trace.Tracer
}

func (s reverser_local_stub) Reverse(ctx context.Context, a0 string) (r0 string, err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "service.Reverser.Reverse", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Reverse(ctx, a0)
}

type reviews_local_stub struct {
	impl   Reviews
	tracer trace.Tracer
}

func (s reviews_local_stub) GetAllBookReviews(ctx context.Context, a0 string) (err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "service.Reviews.GetAllBookReviews", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.GetAllBookReviews(ctx, a0)
}

// Client stub implementations.

type books_client_stub struct {
	stub                   codegen.Stub
	registerBookMetrics    *codegen.MethodMetrics
	findBookByTitleMetrics *codegen.MethodMetrics
}

func (s books_client_stub) RegisterBook(ctx context.Context, a0 Book) (err error) {
	// Update metrics.
	start := time.Now()
	s.registerBookMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "service.Books.RegisterBook", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.registerBookMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.registerBookMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += serviceweaver_size_Book_c456e70f(&a0)
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	(a0).WeaverMarshal(enc)
	var shardKey uint64

	// Call the remote method.
	s.registerBookMetrics.BytesRequest.Put(float64(len(enc.Data())))
	var results []byte
	results, err = s.stub.Run(ctx, 1, enc.Data(), shardKey)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}
	s.registerBookMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

func (s books_client_stub) FindBookByTitle(ctx context.Context, a0 string) (r0 []byte, err error) {
	// Update metrics.
	start := time.Now()
	s.findBookByTitleMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "service.Books.FindBookByTitle", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.findBookByTitleMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.findBookByTitleMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += (4 + len(a0))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.String(a0)
	var shardKey uint64

	// Call the remote method.
	s.findBookByTitleMetrics.BytesRequest.Put(float64(len(enc.Data())))
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}
	s.findBookByTitleMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_slice_byte_87461245(dec)
	err = dec.Error()
	return
}

type details_client_stub struct {
	stub                       codegen.Stub
	addBookDetailsMetrics      *codegen.MethodMetrics
	findDetailsByBookIdMetrics *codegen.MethodMetrics
}

func (s details_client_stub) AddBookDetails(ctx context.Context, a0 string) (err error) {
	// Update metrics.
	start := time.Now()
	s.addBookDetailsMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "service.Details.AddBookDetails", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.addBookDetailsMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.addBookDetailsMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += (4 + len(a0))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.String(a0)
	var shardKey uint64

	// Call the remote method.
	s.addBookDetailsMetrics.BytesRequest.Put(float64(len(enc.Data())))
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}
	s.addBookDetailsMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

func (s details_client_stub) FindDetailsByBookId(ctx context.Context, a0 string) (err error) {
	// Update metrics.
	start := time.Now()
	s.findDetailsByBookIdMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "service.Details.FindDetailsByBookId", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.findDetailsByBookIdMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.findDetailsByBookIdMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += (4 + len(a0))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.String(a0)
	var shardKey uint64

	// Call the remote method.
	s.findDetailsByBookIdMetrics.BytesRequest.Put(float64(len(enc.Data())))
	var results []byte
	results, err = s.stub.Run(ctx, 1, enc.Data(), shardKey)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}
	s.findDetailsByBookIdMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

type reverser_client_stub struct {
	stub           codegen.Stub
	reverseMetrics *codegen.MethodMetrics
}

func (s reverser_client_stub) Reverse(ctx context.Context, a0 string) (r0 string, err error) {
	// Update metrics.
	start := time.Now()
	s.reverseMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "service.Reverser.Reverse", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.reverseMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.reverseMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += (4 + len(a0))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.String(a0)
	var shardKey uint64

	// Call the remote method.
	s.reverseMetrics.BytesRequest.Put(float64(len(enc.Data())))
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}
	s.reverseMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = dec.String()
	err = dec.Error()
	return
}

type reviews_client_stub struct {
	stub                     codegen.Stub
	getAllBookReviewsMetrics *codegen.MethodMetrics
}

func (s reviews_client_stub) GetAllBookReviews(ctx context.Context, a0 string) (err error) {
	// Update metrics.
	start := time.Now()
	s.getAllBookReviewsMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "service.Reviews.GetAllBookReviews", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.getAllBookReviewsMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.getAllBookReviewsMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += (4 + len(a0))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.String(a0)
	var shardKey uint64

	// Call the remote method.
	s.getAllBookReviewsMetrics.BytesRequest.Put(float64(len(enc.Data())))
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}
	s.getAllBookReviewsMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

// Server stub implementations.

type books_server_stub struct {
	impl    Books
	addLoad func(key uint64, load float64)
}

// GetStubFn implements the stub.Server interface.
func (s books_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "RegisterBook":
		return s.registerBook
	case "FindBookByTitle":
		return s.findBookByTitle
	default:
		return nil
	}
}

func (s books_server_stub) registerBook(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 Book
	(&a0).WeaverUnmarshal(dec)

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	appErr := s.impl.RegisterBook(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s books_server_stub) findBookByTitle(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 string
	a0 = dec.String()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.FindBookByTitle(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	serviceweaver_enc_slice_byte_87461245(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

type details_server_stub struct {
	impl    Details
	addLoad func(key uint64, load float64)
}

// GetStubFn implements the stub.Server interface.
func (s details_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "AddBookDetails":
		return s.addBookDetails
	case "FindDetailsByBookId":
		return s.findDetailsByBookId
	default:
		return nil
	}
}

func (s details_server_stub) addBookDetails(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 string
	a0 = dec.String()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	appErr := s.impl.AddBookDetails(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s details_server_stub) findDetailsByBookId(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 string
	a0 = dec.String()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	appErr := s.impl.FindDetailsByBookId(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

type reverser_server_stub struct {
	impl    Reverser
	addLoad func(key uint64, load float64)
}

// GetStubFn implements the stub.Server interface.
func (s reverser_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "Reverse":
		return s.reverse
	default:
		return nil
	}
}

func (s reverser_server_stub) reverse(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 string
	a0 = dec.String()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.Reverse(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.String(r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

type reviews_server_stub struct {
	impl    Reviews
	addLoad func(key uint64, load float64)
}

// GetStubFn implements the stub.Server interface.
func (s reviews_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "GetAllBookReviews":
		return s.getAllBookReviews
	default:
		return nil
	}
}

func (s reviews_server_stub) getAllBookReviews(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 string
	a0 = dec.String()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	appErr := s.impl.GetAllBookReviews(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

// AutoMarshal implementations.

var _ codegen.AutoMarshal = &Book{}

func (x *Book) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("Book.WeaverMarshal: nil receiver"))
	}
	enc.String(x.Title)
	enc.String(x.Author)
	enc.String(x.Description)
}

func (x *Book) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("Book.WeaverUnmarshal: nil receiver"))
	}
	x.Title = dec.String()
	x.Author = dec.String()
	x.Description = dec.String()
}

// Encoding/decoding implementations.

func serviceweaver_enc_slice_byte_87461245(enc *codegen.Encoder, arg []byte) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		enc.Byte(arg[i])
	}
}

func serviceweaver_dec_slice_byte_87461245(dec *codegen.Decoder) []byte {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]byte, n)
	for i := 0; i < n; i++ {
		res[i] = dec.Byte()
	}
	return res
}

// Size implementations.

// serviceweaver_size_Book_c456e70f returns the size (in bytes) of the serialization
// of the provided type.
func serviceweaver_size_Book_c456e70f(x *Book) int {
	size := 0
	size += 0
	size += (4 + len(x.Title))
	size += (4 + len(x.Author))
	size += (4 + len(x.Description))
	return size
}
