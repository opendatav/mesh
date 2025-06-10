/*
 * Copyright (c) 2019, 2024, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

import React from "react";
import {Markdown} from "@/widget";
import readme from "@/router/README.md?raw";
import {Citizen} from "@ducesoft/mesh";
import {TerminalHub} from "@/views";

/**
 * 系统设置中心
 */
export const Administrator: Citizen = {
    name: "系统设置中心",
    path: "/settings/*",
    describe: <Markdown split="系统设置中心">{readme}</Markdown>,
    disable: false,
    route: <div/>,
    children: [{
        name: '命令终端',
        icon: <div></div>,
        path: '/terminal/*',
        route: <div/>,
        disable: true,
        describe: <Markdown split="命令终端">{readme}</Markdown>,
        children: [{
            name: "命令终端",
            path: "/cmd",
            describe: <Markdown split="命令终端">{readme}</Markdown>,
            route: <TerminalHub/>,
        }]
    }, {
        name: '流程管理',
        icon: <div></div>,
        path: '/workflow/*',
        route: <div/>,
        disable: true,
        describe: <Markdown split="流程管理">{readme}</Markdown>,
        children: [{
            name: "流程编排",
            path: "/pipeline",
            describe: <Markdown split="流程编排">{readme}</Markdown>,
            route: <div/>,
        }, {
            name: "流程目录",
            path: "/catalog",
            describe: <Markdown split="流程目录">{readme}</Markdown>,
            route: <div/>,
        }]
    }, {
        name: '其他',
        icon: <div></div>,
        path: '/other/*',
        route: <div/>,
        disable: true,
        describe: <Markdown split="其他">{readme}</Markdown>,
        children: [{
            name: "通知模板管理",
            path: "/notify",
            describe: <Markdown split="通知模板管理">{readme}</Markdown>,
            route: <div/>,
        }, {
            name: "通知事件记录",
            path: "/event",
            describe: <Markdown split="通知事件记录">{readme}</Markdown>,
            route: <div/>,
        }]
    }]
};
