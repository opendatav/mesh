/*
 * Copyright (c) 2019, 2024, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

import '@google/model-viewer';

declare global {

    namespace JSX {

        interface IntrinsicElements {

            'model-viewer': ModelViewer;
        }

        interface ModelViewer {

            'src': string;
            'ios-src'?: string;
            'class'?: string;
            'ar'?: boolean;
            'auto-rotate'?: boolean;
            'camera-controls'?: boolean;
            'shadow-intensity'?: number;
            'shadow-softness'?: number;
            'exposure'?: number;
            'environment-image'?: string;
            'alt'?: string;
            'autoplay'?: boolean;
            'style'?: {};
            'dracoDecoderLocation'?: string;
        }

    }

}
