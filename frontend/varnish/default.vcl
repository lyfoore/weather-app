vcl 4.1;

backend default {
    .host = "backend";
    .port = "8080";
}

sub vcl_recv {
    if (req.method != "GET" && req.method != "HEAD") {
        return (pass);
    }

    return (hash);
}

sub vcl_backend_response {
    set beresp.ttl = 1h;
}