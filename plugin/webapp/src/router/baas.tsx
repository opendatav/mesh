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
 * 基础服务中心
 */
export const BaaS: Citizen = {
    name: "基础服务中心",
    path: "/baas/*",
    describe: <Markdown split="基础服务中心">{readme}</Markdown>,
    disable: false,
    route: <div/>,
    children: [{
        name: '证书管理',
        icon: <div></div>,
        path: '/keys/*',
        route: <div/>,
        disable: true,
        describe: <Markdown split="密钥管理">{readme}</Markdown>,
        children: [{
            name: "许可证书",
            path: "/license",
            describe: <Markdown split="许可证书">{readme}</Markdown>,
            route: <div/>,
        }, {
            name: "通信证书",
            path: "/certificate",
            describe: <Markdown split="通信证书">{readme}</Markdown>,
            route: <div/>,
        }, {
            name: "授权证书",
            path: "/token",
            describe: <Markdown split="授权证书">{readme}</Markdown>,
            route: <div/>,
        }, {
            name: "密钥证书",
            path: "/cipher",
            describe: <Markdown split="密钥证书">{readme}</Markdown>,
            route: <div/>,
        }]
    }, {
        name: '队列管理',
        icon: <div></div>,
        path: '/queue/*',
        route: <div/>,
        disable: true,
        describe: <Markdown split="队列管理">{readme}</Markdown>,
        children: [{
            name: "订阅关系",
            path: "/metadata",
            describe: <Markdown split="订阅关系">{readme}</Markdown>,
            route: <div/>,
        }, {
            name: "消息队列",
            path: "/broker",
            describe: <Markdown split="消息队列">{readme}</Markdown>,
            route: <div/>,
        }]
    }, {
        name: '存储管理',
        icon: <div></div>,
        path: '/store/*',
        route: <div/>,
        disable: true,
        describe: <Markdown split="存储管理">{readme}</Markdown>,
        children: [{
            name: "缓存数据管理",
            path: "/cache",
            describe: <Markdown split="缓存数据管理">{readme}</Markdown>,
            route: <div/>,
        }, {
            name: "关系数据管理",
            path: "/database",
            describe: <Markdown split="关系数据管理">{readme}</Markdown>,
            route: <div/>,
        }, {
            name: "对象数据管理",
            path: "/object",
            describe: <Markdown split="对象数据管理">{readme}</Markdown>,
            route: <div/>,
        }, {
            name: "图数据管理",
            path: "/graph",
            describe: <Markdown split="图数据管理">{readme}</Markdown>,
            route: <div/>,
        }]
    }]
};
