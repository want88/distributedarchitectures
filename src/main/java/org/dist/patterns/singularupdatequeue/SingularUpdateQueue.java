package org.dist.patterns.singularupdatequeue;

import java.util.concurrent.ArrayBlockingQueue;
import java.util.concurrent.CompletableFuture;

public class SingularUpdateQueue<Request, Response> extends Thread {
    private ArrayBlockingQueue<RequestWrapper<Request, Response>> workQueue = new ArrayBlockingQueue<RequestWrapper<Request, Response>>(100);
    private UpdateHandler<Request, Response> updateHandler;
    private SingularUpdateQueue<Response, ?> next;

    public SingularUpdateQueue(UpdateHandler<Request, Response> updateHandler) {
        this.updateHandler = updateHandler;
    }

    public SingularUpdateQueue(UpdateHandler<Request, Response> updateHandler, SingularUpdateQueue<Response, ?> next) {
        this.updateHandler = updateHandler;
        this.next = next;
    }

    static class RequestWrapper<P, R> {
        private final CompletableFuture<R> future;
        private final P request;

        public RequestWrapper(P request) {
            this.request = request;
            this.future = new CompletableFuture<R>();
        }

        public CompletableFuture<R> getFuture() {
            return this.future;
        }

        public P getRequest() {
            return request;
        }

        public void complete(R response) {
            future.complete(response);
        }
    }

    public CompletableFuture<Response> submit(Request request) {
        RequestWrapper<Request, Response> wrapper = new RequestWrapper<>(request);
        workQueue.add(wrapper);
        return wrapper.future;
    }

    @Override
    public void run() {
        try {
            while(true) {
                RequestWrapper<Request, Response> wrapper = workQueue.take();
                Request request = (Request) wrapper.request;
                Response response = updateHandler.update(request);
                if (next != null) {
                    next.submit(response);
                }
                wrapper.future.complete(response);
            }
        } catch (InterruptedException e) {
            throw new RuntimeException(e);
        }
    }
}