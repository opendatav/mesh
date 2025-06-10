/*
 * Copyright (c) 2019, 2024, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

import {useNavigate} from 'react-router-dom';

function Footer(props: any) {

    const navigate = useNavigate();
    const clickMenu = (menu: { name?: string; path: any; }) => {
        navigate(menu.path, {replace: true});
    };

    if ('__TAURI__' in window) {
        return <div/>
    }

    return (
        <div style={{
            display: 'flex',
            justifyContent: 'center',
            alignItems: 'center',
            fontSize: 12,
            margin: '10px auto 10px auto',
            position: 'relative',
            bottom: 0,
            left: 0,
            width: '100%',
            overflowY: 'scroll',
            color: 'var(--color-text-2)'
        }}>
            Copyright © 2023 3344.tech Inc. 保留所有权利。 隐私政策 使用条款 销售政策 法律信息 网站地图
        </div>
    );
}

export default Footer;
