/*
 * Copyright (c) 2019, 2024, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

import {GetCustomToken} from "antd-style";
import {ThemeToken} from "@/types";

export const theme: GetCustomToken<ThemeToken> = (theme) => {
    return {
        x: {
            appbarHeight: '50px',
            appbarIconSize: '40px',
            appbarFontFamily: '-apple-system, BlinkMacSystemFont, "PingFang SC", "Helvetica Neue", Helvetica, Arial, sans-serif',
            buttonAltBG: '#f4f6f7',
            popoverBorderStyle: {
                borderRadius: '0',
                padding: 0,
                border: `1px solid ${theme.token.colorBorder}`
            }
        },
        y: {
            stickHoverStyle: {
                background: `hsla(0, 0%, 100%, 0.9)`,
            }
        }
    }
}
