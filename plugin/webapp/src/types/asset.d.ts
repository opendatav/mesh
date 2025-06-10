/*
 * Copyright (c) 2019, 2024, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

declare module "*.jpg";
declare module "*.woff2";
declare module "*.woff";
declare module "*.ttf";

declare module "*.scss" {
    const content: Record<string, string>;
    export default content;
}

declare module "*.png" {
    const content: Record<string, string>;
    export default content;
}

declare module "*.svg" {
    import React = require('react');
    export const ReactComponent: React.FC<React.SVGProps<SVGSVGElement>>;
    const src: string;
    export default src;
}

declare module "*.md" {
    const content: string;
    export default content;
}

declare module "*.mov" {
    const src: string
    export default src
}

declare module "*.mp4" {
    const src: string
    export default src
}