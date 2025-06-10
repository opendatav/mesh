/*
 * Copyright (c) 2019, 2024, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */


import {Markdown} from "@/widget";
import React from "react";
import readme from "@/router/README.md?raw";
import {Citizen} from "@ducesoft/mesh";
import {DNSCatalog, HTTPCatalog, RouteCatalog} from "@/views";

/**
 * 网络管理中心
 */
export const Network: Citizen = {
    name: "网络管理中心",
    path: "/network/*",
    describe: <Markdown split="网络管理中心">{readme}</Markdown>,
    disable: false,
    route: <div/>,
    children: [{
        name: '网络路由管理',
        icon: <div></div>,
        path: '/route/*',
        route: <div/>,
        disable: true,
        describe: <Markdown split="网络路由管理">{readme}</Markdown>,
        children: [{
            name: "网络路由目录",
            path: "/catalog",
            describe: <Markdown split="网络路由目录">{readme}</Markdown>,
            route: <RouteCatalog/>,
        }, {
            name: "网络路由监控",
            path: "/monitor",
            describe: <Markdown split="网络路由监控">{readme}</Markdown>,
            route: <div/>,
        }]
    }, {
        name: 'HTTP网络管理',
        icon: <div></div>,
        path: '/http/*',
        route: <div/>,
        disable: true,
        describe: <Markdown split="HTTP网络管理">{readme}</Markdown>,
        children: [{
            name: "HTTP服务目录",
            path: "/catalog",
            describe: <Markdown split="HTTP服务目录">{readme}</Markdown>,
            route: <HTTPCatalog/>,
        }, {
            name: "HTTP服务监控",
            path: "/monitor",
            describe: <Markdown split="HTTP服务监控">{readme}</Markdown>,
            route: <div/>,
        }]
    }, {
        name: 'TCP网络管理',
        icon: <div></div>,
        path: '/tcp/*',
        route: <div/>,
        disable: true,
        describe: <Markdown split="安全合规洞察">{readme}</Markdown>,
        children: [{
            name: "TCP服务目录",
            path: "/catalog",
            describe: <Markdown split="TCP服务目录">{readme}</Markdown>,
            route: <div/>,
        }, {
            name: "TCP服务监控",
            path: "/monitor",
            describe: <Markdown split="TCP服务监控">{readme}</Markdown>,
            route: <div/>,
        }]
    }, {
        name: 'UDP网络管理',
        icon: <div></div>,
        path: '/udp/*',
        route: <div/>,
        disable: true,
        describe: <Markdown split="UDP网络管理">{readme}</Markdown>,
        children: [{
            name: "UDP服务目录",
            path: "/catalog",
            describe: <Markdown split="UDP服务目录">{readme}</Markdown>,
            route: <div/>,
        }, {
            name: "UDP服务监控",
            path: "/monitor",
            describe: <Markdown split="UDP服务监控">{readme}</Markdown>,
            route: <div/>,
        }]
    }, {
        name: 'DNS网络管理',
        icon: <div></div>,
        path: '/dns/*',
        route: <div/>,
        disable: true,
        describe: <Markdown split="DNS网络管理">{readme}</Markdown>,
        children: [{
            name: "DNS服务目录",
            path: "/catalog",
            describe: <Markdown split="DNS服务目录">{readme}</Markdown>,
            route: <DNSCatalog/>,
        }, {
            name: "DNS服务监控",
            path: "/monitor",
            describe: <Markdown split="DNS服务监控">{readme}</Markdown>,
            route: <div/>,
        }]
    }]
};
