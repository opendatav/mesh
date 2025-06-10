/*
 * Copyright (c) 2000, 2024, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

import type {ProColumns as Columns} from '@ant-design/pro-components';
import {ProTable} from "@ant-design/pro-components";
import {ExportOutlined} from '@ant-design/icons';
import {Button, Tag} from 'antd';
import {Codec, context, Mesh, MeshService, Page, Paging, ServiceLoader, States} from '@ducesoft/mesh';
import {Link} from 'react-router-dom';
import moment from "moment";
import {DNSCreator} from "@/views";
import {useState} from "react";

const columns: Columns<Record<string, any>>[] = [
    {
        title: '名称',
        dataIndex: 'name',
        ellipsis: true,
    }, {
        title: '类型',
        dataIndex: 'kind',
        ellipsis: true,
        search: false,
    }, {
        title: '规则',
        dataIndex: 'pattern',
        ellipsis: true,
        search: false,

    }, {
        title: '地址',
        dataIndex: 'address',
        ellipsis: true,
        search: false,
    }, {
        title: '上行带宽',
        dataIndex: 'upstream',
        ellipsis: true,
        search: false,
        render: (_, row) => {
            return `${row?.upstream}Mbps`;
        }
    }, {
        title: '下行带宽',
        dataIndex: 'downstream',
        ellipsis: true,
        search: false,
        render: (_, row) => {
            return `${row?.downstream}Mbps`;
        }
    }, {
        title: '节点编号',
        dataIndex: 'node_id',
        ellipsis: true,
        search: false,
    }, {
        title: '机构编号',
        dataIndex: 'inst_id',
        ellipsis: true,
        search: false,
    }, {
        disable: true,
        title: '状态',
        dataIndex: 'status',
        filters: true,
        onFilter: true,
        ellipsis: true,
        valueType: 'select',
        valueEnum: States.Route.asMap(),
        render: (_, row) => {
            const x = States.Route.from(row?.status);
            return <Tag color={x.color}>{x.text}</Tag>;
        }
    }, {
        title: '创建时间',
        dataIndex: 'createAt',
        ellipsis: true,
        search: false,
        render: (_, row) => {
            return moment(row?.createAt).format('YYYY-MM-DD HH:mm:ss');
        }
    }, {
        title: '操作',
        valueType: 'option',
        key: 'option',
        render: (dom, row, idx, action) => {
            if (row?.status == 32) {
                return [<Link to={`describe?id=${row.id}`}>查看</Link>,
                    <a href='#' rel="noopener noreferrer" key="view">启用</a>]
            }
            return [<Link to={`describe?id=${row.id}`}>查看</Link>,
                <a href='#' rel="noopener noreferrer" key="view">禁用</a>]
        }
    }];

export function VPNCatalog() {

    const codec = ServiceLoader.load(Codec).getDefault();
    const [open, setOpen] = useState(false);

    return (
        <div>
            <ProTable<Record<string, any>>
                columns={columns}
                cardBordered
                request={async (idx, sort, filter) => {
                    const ctx = context();
                    ctx.setAttribute(Mesh.UNAME, "mesh.dns.index");
                    const paging: Paging = {
                        sid: '',
                        index: idx.current || 1 - 1,
                        limit: idx.pageSize || 10,
                        factor: {},
                        order: ''
                    };
                    const raw = await MeshService.endpoint.fuzzy(codec.encode(paging), ctx);
                    return codec.decode(raw, Object) as Page<Record<string, any>>;
                }}
                rowKey="id"
                editable={{
                    type: 'multiple',
                }}
                search={{
                    labelWidth: 'auto',
                }}
                pagination={{
                    pageSize: 10,
                }}
                dateFormatter="string"
                toolBarRender={() => [
                    <Button key="button" icon={<ExportOutlined/>} type="primary" onClick={() => setOpen(true)}>
                        新增
                    </Button>,
                ]}
            />
            <DNSCreator open={open} onClose={() => setOpen(false)}/>
        </div>
    );
}
