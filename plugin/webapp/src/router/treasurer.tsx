/*
 * Copyright (c) 2019, 2024, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */


import React from "react";
import {Markdown} from "@/widget";
import readme from "@/router/README.md?raw";
import {Roadmap} from "@/views";
import {Citizen} from "@ducesoft/mesh";

/**
 * 快捷指引
 */
export const Treasurer: Citizen = {
    name: "快捷指引",
    path: "/treasurer/*",
    describe: <Markdown split="快捷指引">{readme}</Markdown>,
    disable: false,
    route: <div/>,
    children: [{
        name: '迭代计划',
        icon: <div></div>,
        path: '/plan/*',
        route: <div/>,
        disable: true,
        describe: <Markdown split="迭代计划">{readme}</Markdown>,
        children: [{
            name: "路线图",
            path: "/roadmap",
            describe: <Markdown split="路线图">{readme}</Markdown>,
            route: <Roadmap/>,
        }]
    }]
};
