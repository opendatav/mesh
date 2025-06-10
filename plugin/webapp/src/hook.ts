/*
 * Copyright (c) 2019, 2024, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

import {Hook} from '@ducesoft/mesh';

Hook.onFailure('eas.hook', (ctx, invocation, c) => {
    if ('E0000000609' == c.getCode()) {
        window.location.href = '/x/license';
        return
    }
    if ('E0000000630' == c.getCode()) {
        // window.location.href = '/x/authorize';
        return
    }
    // Message.error(c.getMessage());
});