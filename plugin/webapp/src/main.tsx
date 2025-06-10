/*
 * Copyright (c) 2019, 2024, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

import React from 'react';
import ReactDOM from 'react-dom/client';
import {BrowserRouter} from "react-router-dom";
import RunC from '@/runc';
import '@/main.css';

const root = ReactDOM.createRoot(document.getElementById('root')!);

root.render(
    <BrowserRouter>
        <RunC/>
    </BrowserRouter>
);


