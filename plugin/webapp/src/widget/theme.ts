/*
 * Copyright (c) 2019, 2024, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

import {createStyles} from "antd-style";

export const useStyles = createStyles(({token, css}) => ({
    scrollBar: css`
        ::-webkit-scrollbar {
            width: 2px;
            height: 2px;
        }

        ::-webkit-scrollbar-track {
            background-color: #f1f1f1; /* 滚动条轨道颜色 */
        }

        ::-webkit-scrollbar-thumb {
            background-color: rgba(0, 111, 255, 1); /* 滚动条滑块颜色 */
            border-radius: 1px; /* 滑块圆角 */
        }

        ::-webkit-scrollbar-thumb:hover {
            background-color: grey; /* 滑块悬停时的颜色 */
        }
    `,
    collapseView: css`
        padding: 0 16px;
        height: 48px;
        line-height: 48px;
        border-top: 1px solid #ebebeb;
        border-bottom: 1px solid #ebebeb;
        width: 100%;
        max-width: 100%;
        cursor: pointer;
        vertical-align: middle;
        text-align: left;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
        color: #666;
        background-color: rgba(0, 46, 70, 0.04314);
        font-size: 12px;
    `,
    leftView: css`
        font-size: 12px;
        will-change: left;
        width: 250px;
        height: 100%;
        background-color: #fff;
        box-shadow: 1px 0 2px 0 rgba(0, 0, 0, 0.16);
        transform: translateX(0px);
        overflow: hidden;
        transition: transform 124ms ease-in-out, left 62ms;
    `,
    subsystemView: css`
        width: 100%;
        max-width: 100%;
        cursor: pointer;
        vertical-align: middle;
        text-align: left;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
        background-color: transparent;
        padding: 0 32px 0 42px;
        height: 32px;
        line-height: 30px;
        font-size: 12px;
        transition: 0.3s ease-out;
        color: rgba(6, 6, 6, 0.5);

        :hover {
            color: #006fff;
            border-right: 2px solid #006fff;
        }
    `,
    rightView: css`
        padding: 0 16px 0 16px;
        width: 100%;
        overflow: scroll;
    `,
    subsystemTitle: css`
        height: 32px;
        font-size: 14px;
        font-weight: 600;
        line-height: 32px;
        display: inline-flex;
        width: 100%;
        -webkit-box-align: center;
        align-items: center;
        text-overflow: ellipsis;
        white-space: nowrap;
        color: #333;

        :after {
            content: "";
            flex: 1 1 0;
            border-right: none;
            border-bottom: none;
            border-left: none;
            border-image: initial;
            margin-left: 0.5em;
            border-top: 1px solid #ebebeb;
        }
    `,
    plateTitle: css`
        height: 32px;
        font-size: 12px;
        font-weight: 600;
        line-height: 32px;
        display: inline-flex;
        width: 100%;
        -webkit-box-align: center;
        align-items: center;
        text-overflow: ellipsis;
        white-space: nowrap;
        color: #333;
        padding: 0 0 0 5px;
    `,
    itemView: css`
        width: 100%;
        max-width: 100%;
        cursor: pointer;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
        height: 32px;
        line-height: 30px;
        font-size: 12px;
        font-family: -apple-system, "system-ui", "PingFang SC", "Helvetica Neue", Helvetica, Arial, sans-serif;
        transition: 0.3s ease-out;
        color: rgba(6, 6, 6, 0.5);
        padding: 0 0 0 5px;

        :hover {
            color: #006fff;
            background: rgba(0, 111, 255, 0.1);
        }
    `,
    buttonView: css`
        display: inline-block;
        border: 1px solid transparent;
        border-radius: 2px;
        width: auto;
        max-width: 100%;
        cursor: pointer;
        vertical-align: middle;
        text-align: center;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
        margin: 0 12px;
        color: #666;
        background-color: ${token.x.buttonAltBG};
        padding: 0 12px;
        height: 32px;
        line-height: 30px;
        font-size: 12px;
        transition: 0.3s ease-out;
    `,
    sideMenuView: css`
        overflow: hidden auto;
        order: 0;
        flex: 0 1 auto;
        box-sizing: content-box;
        width: 208px;
        height: 100%;
        background-color: rgb(255, 255, 255);
        border-right: 1px solid rgb(203, 203, 203);
        min-width: 208px;
        padding: 0 0 0 20px;
        transition: all 1.5s ease-out;
    `,
    sideMenuTitle: css`
        background-color: rgb(255, 255, 255);
        position: relative;
        font-size: 14px;
        font-weight: 600;
        color: rgb(51, 51, 51);
        line-height: 18px;
        box-sizing: border-box;
        padding: 24px 0 10px 0;
        width: 100%;
    `,
    sideMenuGroup: css`
        place-content: stretch space-between;
        min-height: 32px;
        list-style: none;
        cursor: pointer;
        font-weight: 700;
        color: rgb(51, 51, 51);
        line-height: 20px;
        font-size: 12px;
        padding: 10px 0 0 5px;
    `,
    //
    stickyView: css`
        top: 0;
        right: 0;
        left: 0;
        position: sticky;
        z-index: 1000;
        height: ${token.x.appbarHeight};
        line-height: 1.5;
        font-size: 12px;
        font-family: ${token.x.appbarFontFamily}, -apple-system, BlinkMacSystemFont, "PingFang SC", "Helvetica Neue", Helvetica, Arial, sans-serif;
        border-bottom: 1px solid rgba(0, 0, 0, 0.1);
        box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    `,
    linkButton: css`
        display: inline-block;
        border: 1px solid transparent;
        border-radius: 2px;
        width: auto;
        max-width: 100%;
        cursor: pointer;
        vertical-align: middle;
        text-align: center;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
        margin: 0 12px;
        color: #666;
        background-color: ${token.x.buttonAltBG};
        padding: 0 12px;
        height: 32px;
        line-height: 30px;
        font-size: 12px;
        transition: 0.3s ease-out;
    `,
    iconButton: css`
        font-size: 17px;
        color: #666;
        margin: 0 14px;
        cursor: pointer;
    `,
    menuButton: css`
        position: relative;
        width: ${token.x.appbarHeight};
        height: ${token.x.appbarHeight};
        background-color: #006fff;
        border: 0 solid transparent;
        margin-right: 1rem;
        color: #fff;
        font-size: 1.5rem;
    `,
    buttonIcon: css`
        display: block;
        background-color: rgb(255, 255, 255);
        width: 20px;
        height: 2px;
        margin: 0.15rem 0;
        transition: 0.3s ease-out;
        transform-origin: left center;
    `,
    drawerView: css`
        position: relative;
        overflow: hidden;
        height: calc(100vh - ${token.x.appbarHeight});
    `,
    nameText: css`
        font-size: 0.9em;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
        color: #999;
    `,
    logonIdText: css`
        line-height: 20px;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
        color: #333;
    `,
    linkText: css`
        font-size: 0.9em;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
        color: #006fff;
    `,
    padding20: css`
        padding: 10px 20px;
    `
}));
