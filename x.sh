curl -X POST http://127.0.0.1:7305/mpc/invoke -H "Content-Type: application/json" -d '
{
    "method": "mesh.net.route.add",
    "content": {
        "routes": [
            {
                "name": "测试路由",
                "listen": [
                    "transport_x"
                ],
                "matcher": "ClientIP(`x.x.x.x`) || Method(`POST`) || Host(`x.com`) || HostRegexp(`x.*`) || Path(`/a/b/c`) || PathRegexp(`/a/b/c/.*`) || Header(`x-header`, `x-value`) || HeaderRegexp(`x-.*`) || Query(`x-query`, `x-value`) || QueryRegexp(`x-query`, `x.*`)",
                "backend": [
                    "http://x.x.x.x:port"
                ],
                "priority": 1,
                "pass_host_header": true
            }
        ]
    }
}
'

curl -X POST http://127.0.0.1:7305/mpc/invoke -H "Content-Type: application/json" -d '
{
    "method": "mesh.net.route.remove",
    "content": {
        "names": ["测试路由"]
    }
}
'