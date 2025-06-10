/*
 * Copyright (c) 2019, 2024, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

import {Route, Routes} from "react-router-dom";
import {App, ConfigProvider} from "antd";
import {StyleProvider, ThemeProvider} from "antd-style";
import {MeshHub} from "@/router";
import {theme} from "@/theme";
import '@/hook';
import '@/runc.css';
import 'antd/dist/reset.css';

function RunC() {

    return (
        <ConfigProvider>
            <ThemeProvider defaultThemeMode={'light'} customToken={theme}>
                <StyleProvider>
                    <App>
                        <Routes>
                            <Route path="/mesh/*" element={<MeshHub/>}/>
                            <Route path="/*" element={<MeshHub/>}/>
                        </Routes>
                    </App>
                </StyleProvider>
            </ThemeProvider>
        </ConfigProvider>
    );
}

export default RunC;