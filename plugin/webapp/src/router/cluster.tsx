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

/**
 * 集群管理中心
 */
export const Cluster: Citizen = {
    name: "集群管理中心",
    path: "/cluster/*",
    describe: <Markdown split="集群管理中心">{readme}</Markdown>,
    disable: false,
    route: <div/>,
    children: [{
        name: '集群管理',
        icon: <div></div>,
        path: '/cluster/*',
        route: <div/>,
        disable: true,
        describe: <Markdown split="集群管理">{readme}</Markdown>,
        children: [{
            name: "集群目录",
            path: "/catalog",
            describe: <Markdown split="集群目录">{readme}</Markdown>,
            route: <div/>,
        }, {
            name: "集群监控",
            path: "/monitor",
            describe: <Markdown split="集群监控">{readme}</Markdown>,
            route: <div/>,
        }]
    }, {
        name: '服务治理',
        icon: <div></div>,
        path: '/service/*',
        route: <div/>,
        disable: true,
        describe: <Markdown split="服务治理">{readme}</Markdown>,
        children: [{
            name: "服务目录",
            path: "/catalog",
            describe: <Markdown split="服务目录">{readme}</Markdown>,
            route: <div/>,
        }, {
            name: "服务监控",
            path: "/monitor",
            describe: <Markdown split="服务监控">{readme}</Markdown>,
            route: <div/>,
        }]
    }, {
        name: '配置管理',
        icon: <div></div>,
        path: '/configmap/*',
        route: <div/>,
        disable: true,
        describe: <Markdown split="配置管理">{readme}</Markdown>,
        children: [{
            name: "配置目录",
            path: "/catalog",
            describe: <Markdown split="配置目录">{readme}</Markdown>,
            route: <div/>,
        }, {
            name: "推送历史",
            path: "/history",
            describe: <Markdown split="推送历史">{readme}</Markdown>,
            route: <div/>,
        }]
    }, {
        name: '服务地图',
        icon: <div></div>,
        path: '/map/*',
        route: <div/>,
        disable: true,
        describe: <Markdown split="服务地图">{readme}</Markdown>,
        children: [{
            name: "依赖图谱",
            path: "/depends",
            describe: <Markdown split="依赖图谱">{readme}</Markdown>,
            route: <div/>,
        }, {
            name: "调用图谱",
            path: "/trace",
            describe: <Markdown split="调用图谱">{readme}</Markdown>,
            route: <div/>,
        }, {
            name: "日志查询",
            path: "/tail",
            describe: <Markdown split="日志查询">{readme}</Markdown>,
            route: <div/>,
        }]
    }]
};
