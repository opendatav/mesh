/*
 * Copyright (c) 2019, 2024, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

export async function readFile(f: File): Promise<Uint8Array> {
    return new Promise((resolve, reject) => {
        const reader = new FileReader();
        reader.readAsArrayBuffer(f);
        reader.onload = function (e) {
            resolve(new Uint8Array(e?.target?.result as ArrayBuffer));
        }
        reader.onerror = function (e) {
            reject(e);
        }
        reader.onabort = function (e) {
            reject(e);
        }
    })
}