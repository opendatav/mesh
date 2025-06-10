/*
 * Copyright (c) 2019, 2024, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

declare interface Window {
    __TAURI__?: {
        writeText(text: string): Promise<void>;
    };
    ModelViewerElement?: {
        dracoDecoderLocation?: string;
    };
}
