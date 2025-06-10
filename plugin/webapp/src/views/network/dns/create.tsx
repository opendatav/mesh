/*
 * Copyright (c) 2000, 2024, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

import {
    ModalForm,
    ProForm,
    ProFormRadio,
    ProFormSelect,
    ProFormText,
    ProFormTextArea
} from '@ant-design/pro-components';
import {Codec, context, Mesh, MeshService, ServiceLoader} from '@ducesoft/mesh';
import {Form} from 'antd';
import React, {useEffect, useState} from 'react';

export function DNSCreator(props: { open: boolean, onClose?: () => void }) {

    const codec = ServiceLoader.load(Codec).getDefault();
    const [form] = Form.useForm<{ name: string; company: string }>();
    const [open, setOpen] = useState(props.open);

    useEffect(() => {
        setOpen(props.open);
    }, [props]);

    const onUpsert = async (dict: Record<string, any>) => {
        const ctx = context();
        ctx.setAttribute(Mesh.UNAME, "mesh.dns.upsert");
        const idx = {
            id: '10000',
            zone: 'cn-hangzhou',
            ttl: 120,
            tags: [],
            raw: '{}',
            provider: '',
            pid: '',
            memo: '',
        };
        await MeshService.endpoint.fuzzy(codec.encode(Object.assign(idx, dict)), ctx);
    }

    return (
        <ModalForm<{ name: string; company: string; }>
            title="新增DNS解析记录"
            open={open}
            form={form}
            autoFocusFirstInput
            modalProps={{destroyOnClose: true,}}
            submitTimeout={2000}
            onFinish={async (values) => {
                await onUpsert(values);
                props?.onClose?.();
                return true;
            }}
            onOpenChange={(open) => {
                if (!open) {
                    props?.onClose?.();
                }
            }}
        >
            <ProForm.Group>
                <ProFormText
                    width="md"
                    name="domain"
                    label="域名"
                    placeholder="请输入域名"
                />
                <ProFormRadio.Group
                    name="kind"
                    label="类型"
                    options={[{label: 'A', value: 'A'}, {label: 'AAAA', value: 'AAAA'}]}
                />
            </ProForm.Group>
            <ProForm.Group>
                <ProFormText
                    width="md"
                    name="text"
                    label="记录"
                    tooltip="最长为32位"
                    placeholder="请输入解析记录"
                />
                <ProFormRadio.Group
                    name="provider"
                    label="解析方式"
                    options={[{label: '动态解析', value: 'DYN'}, {label: '静态解析', value: 'FIX'}]}
                />
            </ProForm.Group>
            <ProForm.Group>
                <ProFormSelect name="provider" label="解析服务提供商" width="md" valueEnum={{
                    alibaba: '阿里云解析',
                    mesh: '网格内置解析',
                }} placeholder="如：阿里云解析" rules={[{required: true}]}
                />
                <ProFormSelect
                    width="md"
                    name="lables"
                    label="标签"
                    mode="tags"
                    placeholder="请输入标签"
                />
            </ProForm.Group>
            <ProFormTextArea
                name="memo"
                label="详细说明"
                placeholder="请输入备注"
            />
        </ModalForm>
    )
}
