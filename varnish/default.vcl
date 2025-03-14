vcl 4.1;

backend api {
    .host = "backend";
    .port = "8080";
}

backend static {
    .host = "frontend";
    .port = "80";
}

sub vcl_recv {
    if (req.url ~ "^/api/") {
        set req.backend_hint = api;
    } else {
        set req.backend_hint = static;
    }

    if (req.method != "GET" && req.method != "HEAD") {
        return (pass);
    }

    if (req.url ~ "\.(css|js|png|jpg|gif|ico|webp)$") {
        unset req.http.Cookie;
    }

    return (hash);
}

sub vcl_backend_response {
    if (bereq.url ~ "^/api/") {
        set beresp.ttl = 30m;
    } else {
        set beresp.ttl = 1d;
    }
}

sub vcl_deliver {
    if (obj.hits > 0) {
        set resp.http.X-Cache = "HIT";
    } else {
        set resp.http.X-Cache = "MISS";
    }
}