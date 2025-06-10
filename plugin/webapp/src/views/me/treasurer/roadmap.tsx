/*
 * Copyright (c) 2019, 2024, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

import React from 'react';
import {Divider, Typography} from 'antd';

const {Title, Paragraph, Text, Link} = Typography;

export function Roadmap() {

    return (
        <div style={{overflow: 'scroll', height: '80vh'}}>
            <Typography>
                <Title>摘要</Title>
                <Paragraph>
                    冷链平台一期主要讲三个生态故事线：
                    <ul>
                        <li>
                            <Link href="#">技术服务商数据生态：冷链链接玩转数据技术服务商（蓝象、隐语、FATE等）</Link>
                        </li>
                        <li>
                            <Link href="#">平台服务商数据生态：冷链链接玩转平台服务商（确权、流通、交易、入表等）</Link>
                        </li>
                        <li>
                            <Link
                                href="#">KA商家数据供需生态：冷链链接玩转KA数据商家（针对数据供给方、需求方进行数据供需链接驱动）</Link>
                        </li>
                    </ul>
                    三个商务BD方向：
                    <ul>
                        <li>
                            <Link href="#">国家18个试点项目：冷链平台</Link>
                        </li>
                        <li>
                            <Link href="#">省地市数据主管单位：城市数据空间</Link>
                        </li>
                        <li>
                            <Link href="#">数据领域专精小单品：9个子系统拆分独立售卖</Link>
                        </li>
                    </ul>
                </Paragraph>
                <Paragraph>
                    冷链平台的目标是构建多套标准的数据服务能力集，链接数据生态系统中的各类服务商（
                    <Text strong>
                        硬件技术、软件技术、平台服务、咨询服务、数据供给、数据流转、数据使用
                    </Text>
                    等）。让这些服务商提供的服务/能力能统一、快速、高效、简单、安全的在全国形成网络化、数字化、智能化的生产工具。构建新一代的数据生产关系。
                </Paragraph>
                <Title level={2}>一期版本计划</Title>
                <Divider/>
                <Paragraph>
                    <Text strong>1008版本</Text>（10月8日部署到UAT环境提供对客演示）
                    <br/>
                    故事线：主讲故事线一（技术服务商数据生态）
                    <br/>
                    目标：组合蓝象产品阵列进行演示，跑通GAIA、Glite、GaiaC与冷链平台的链路，需要实现以下板块：
                    <br/>
                    <ul>
                        <li>
                            <Link href="#">冷链开放平台：门户</Link>
                        </li>
                        <li>
                            <Link href="#">冷链开放平台：工作台</Link>
                        </li>
                        <li>
                            <Link href="#">冷链管理平台：1+X纳管中心</Link>
                        </li>
                        <li>
                            <Link href="#">冷链管理平台：冷链网络中心</Link>
                        </li>
                        <li>
                            <Link href="#">冷链管理平台：冷链数据目录：数据资源目录</Link>
                        </li>
                        <li>
                            <Link href="#">冷链管理平台：冷链数据目录：数据流通目录</Link>
                        </li>
                        <li>
                            <Link href="#">冷链管理平台：使用控制中心：使用策略中心</Link>
                        </li>
                        <li>
                            <Link href="#">冷链管理平台：使用控制中心：策略执行过程</Link>
                        </li>
                        <li>
                            <Link href="#">冷链管理平台：冷链开放中心：注册认证管理</Link>
                        </li>
                        <li>
                            <Link href="#">冷链管理平台：冷链开放中心：开放应用管理</Link>
                        </li>
                        <li>
                            <Link href="#">冷链管理平台：冷链监管中心：冷链大屏</Link>
                        </li>
                    </ul>
                </Paragraph>
                <Divider/>
                <Paragraph>
                    <Text strong>1022版本</Text>（10月22日部署到UAT环境提供对客演示）
                    <br/>
                    故事线：主讲故事线一（平台服务商数据生态）
                    <br/>
                    目标：组合数据登记确权平台、试衣间、资产入表等进行演示，需要实现以下板块，目标跑通数据冷链平台服务流转支撑：
                    <br/>
                    <ul>
                        <li>
                            <Link href="#">冷链管理平台：冷链开放中心：所有功能</Link>
                        </li>
                        <li>
                            <Link href="#">冷链管理平台：使用控制中心：所有功能</Link>
                        </li>
                        <li>
                            <Link href="#">冷链管理平台：冷链数据目录：所有功能</Link>
                        </li>
                    </ul>
                </Paragraph>
                <Divider/>
                <Paragraph>
                    <Text strong>1031版本</Text>（10月31日部署到UAT环境提供对客演示，且可对外部署）
                    <br/>
                    故事线：主讲故事线三（KA商家数据供需生态）、监管管控
                    <br/>
                    目标：以冷链应用生态实现数据供需快速建立和流转，极低成本快速完成数据供需链路的仓配管理；监管闭环
                    <br/>
                    <ul>
                        <li>
                            <Link href="#">冷链管理平台：1+X纳管中心：新增多个生态应用</Link>
                        </li>
                        <li>
                            <Link href="#">冷链管理平台：冷链算力中心：所有功能</Link>
                        </li>
                        <li>
                            <Link href="#">冷链管理平台：冷链监控中心：所有功能</Link>
                        </li>
                        <li>
                            <Link href="#">冷链管理平台：冷链监管中心：所有功能</Link>
                        </li>
                        <li>
                            <Link href="#">冷链管理平台：冷链审计中心：所有功能</Link>
                        </li>
                        <li>
                            <Link href="#">冷链管理平台：冷链系统设置：半中心化架构信息管理</Link>
                        </li>
                    </ul>
                </Paragraph>
                <Divider/>
            </Typography>
        </div>
    )
}
