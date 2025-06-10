/*
 * Copyright (c) 2000, 2024, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

import type {ActionType, ProColumns as Columns} from '@ant-design/pro-components';
import {ProTable} from "@ant-design/pro-components";
import {ExportOutlined} from '@ant-design/icons';
import {Button, Popover, Tag} from 'antd';
import {Codec, context, Mesh, MeshService, Paging, ServiceLoader, VTable} from '@ducesoft/mesh';
import {useEffect, useRef, useState} from "react";
import {ServiceDescriber} from "@/views";

const columns: Columns<Record<string, any>>[] = [
    {
        title: '提供方',
        dataIndex: 'provider',
        ellipsis: true,
        search: false,
        width: 100,
    }, {
        title: '入口点',
        dataIndex: 'entrypoint',
        ellipsis: true,
        width: 100,
        render: (_, row) => {
            return (
                <div>
                    {
                        (row.entryPoints as string[]).map(x => {
                            return <Tag key={x} color="blue">{x?.toUpperCase()}</Tag>
                        })
                    }
                </div>
            )
        }
    }, {
        title: '名称',
        dataIndex: 'name',
        ellipsis: true,
        search: false,
        width: 380,
    }, {
        title: '优先级',
        dataIndex: 'priority',
        ellipsis: true,
        search: false,
        width: 100,
        render: (_, row) => {
            if (row.priority > 99999) {
                return <Tag color="green">99999</Tag>
            }
            return <Tag color="green">{row.priority}</Tag>
        }
    }, {
        title: '匹配模式',
        dataIndex: 'rule',
        ellipsis: true,
        search: false,
    }, {
        title: '状态',
        dataIndex: 'status',
        ellipsis: true,
        search: false,
        width: 100,
        render: (_, row) => {
            return (
                <Tag color={row.status == 'enabled' ? 'blue' : 'red'}>
                    {row.status == 'enabled' ? '可用' : '不可用'}
                </Tag>);
        }
    }, {
        title: '操作',
        valueType: 'option',
        key: 'option',
        width: 100,
        render: (dom, row, idx, action) => {
            return (
                <Popover placement={'left'} content={<ServiceDescriber service={row.services || {}}/>}>
                    <span style={{cursor: 'pointer', color: 'deepskyblue'}}>服务详情</span>
                </Popover>
            )
        }
    }];

export function HTTPCatalog() {

    const ref = useRef<ActionType>(null);
    const codec = ServiceLoader.load(Codec).getDefault();
    const [table, setTable] = useState<VTable<Record<string, any>>>(new VTable<Record<string, any>>([]));

    useEffect(() => {
        onLoad().then(table => {
            setTable(table);
            ref?.current?.reload(true);
        });
    }, []);

    const onLoad = async () => {
        const services = await onServiceLoad();
        const routes = await onRouteLoad();
        for (let route of routes) {
            route.services = services[`${route.service}@${route.provider}`] || services[route.service] || {};
        }
        return new VTable<Record<string, any>>(routes);
    }

    const onRouteLoad = async () => {
        const ctx = context();
        ctx.setAttribute(Mesh.UNAME, "mesh.dot.routes");
        const paging: Paging = {
            sid: '',
            index: 0,
            limit: 10000,
            factor: {kind: 'http'},
            order: ''
        };
        const raw = await MeshService.endpoint.fuzzy(codec.encode(paging), ctx);
        return codec.decode(raw, Object) as Record<string, any>[];
    }

    const onServiceLoad = async () => {
        const ctx = context();
        ctx.setAttribute(Mesh.UNAME, "mesh.dot.services");
        const paging: Paging = {
            sid: '',
            index: 0,
            limit: 10000,
            factor: {kind: 'http'},
            order: ''
        };
        const raw = await MeshService.endpoint.fuzzy(codec.encode(paging), ctx);
        const dict: Record<string, Record<string, any>> = {};
        for (let item of (codec.decode(raw, Object) as Record<string, any>[])) {
            dict[item.name] = item;
        }
        return dict;
    }

    return (
        <ProTable<Record<string, any>>
            columns={columns}
            cardBordered
            actionRef={ref}
            request={async (idx, sort, filter) => {
                return table.index({
                    sid: '',
                    index: (idx.current || 1) - 1,
                    limit: idx.pageSize || 10,
                    factor: {
                        name: idx.keyword || ''
                    },
                    order: ''
                });
            }}
            editable={{type: 'multiple',}}
            rowKey="name"
            search={{labelWidth: 'auto',}}
            pagination={{pageSize: 10,}}
            dateFormatter="string"
            toolBarRender={() => [
                <Button key="button" icon={<ExportOutlined/>} type="primary">
                    导出
                </Button>,
            ]}
        />
    );
}
