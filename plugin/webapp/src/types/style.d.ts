/*
 * Copyright (c) 2019, 2024, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

import {CSSProperties} from "react";


export interface ThemeToken {
    x: {
        appbarHeight: string;
        appbarIconSize: string;
        appbarFontFamily: string;
        buttonAltBG: string;
        popoverBorderStyle: CSSProperties;
    };
    y: {
        stickHoverStyle: CSSProperties;
    };
}

declare module 'antd-style' {
    // eslint-disable-next-line @typescript-eslint/no-empty-interface
    export interface CustomToken extends ThemeToken {
    }
}
