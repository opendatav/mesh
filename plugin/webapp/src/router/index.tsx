/*
 * Copyright (c) 2019, 2024, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

import React, {useEffect, useState} from "react";
import {Route, Routes} from "react-router-dom";
import {Citizen, Environ} from "@ducesoft/mesh";
import {Breadcrumb, Flex, Tooltip, Watermark} from "antd";
import {AppstoreOutlined, QuestionCircleOutlined, RubyOutlined} from "@ant-design/icons";
import {createStyles} from "antd-style";
import {Container, SideMenu} from "@/widget";
import {Administrator} from "@/router/administrator";
import {Treasurer} from "@/router/treasurer";
import {Cluster} from "@/router/cluster";
import {Network} from "@/router/network";
import {BaaS} from "@/router/baas";

const useStyles = createStyles(({token, css}) => ({
    drawerView: css`
        position: relative;
        overflow: hidden;
        height: calc(100vh - ${token.x.appbarHeight});
    `,
    crumbView: css`
        padding: 1.5rem 0;
    `
}));

export function MeshHub() {

    const {styles, cx, theme} = useStyles();
    const [environ, setEnviron] = useState(new Environ());
    const [arches, setArches] = useState<Citizen[]>([]);
    const [pivots, setPivots] = useState<Citizen[]>([]);
    const [subsystem, setSubsystem] = useState<Citizen>(new Citizen());
    const [crumbs, setCrumbs] = useState<Citizen[]>([]);

    const staticArches = [{
        name: "服务网格子系统",
        icon: <AppstoreOutlined/>,
        path: "/subsystem/*",
        describe: <div/>,
        children: [Cluster, Network, BaaS, Administrator]
    }, {
        name: "我的收藏",
        icon: <RubyOutlined/>,
        path: "/me/*",
        describe: <div/>,
        children: [Treasurer, {
            name: "常用功能",
            path: "/favorite/*",
            describe: <div/>,
            route: <div>常用功能</div>
        }]
    }];
    const staticPivots = [{
        name: "华北",
        path: "#",
        describe: "",
        children: [{
            name: "北京-北京",
            path: "#",
            describe: ""
        }, {
            name: "天津-天津",
            path: "#",
            describe: ""
        }, {
            name: "山东-青岛",
            path: "#",
            describe: ""
        }]
    }, {
        name: "华中",
        path: "#",
        describe: "",
        children: [{
            name: "湖北-武汉",
            path: "#",
            describe: ""
        }, {
            name: "安徽-合肥",
            path: "#",
            describe: ""
        }, {
            name: "河南-郑州",
            path: "#",
            describe: ""
        }]
    }, {
        name: "华南",
        path: "#",
        describe: "",
        children: [{
            name: "福建-福州",
            path: "#",
            describe: ""
        }, {
            name: "广东-深圳",
            path: "#",
            describe: ""
        }]
    }, {
        name: "西南",
        path: "#",
        describe: "",
        children: [{
            name: "重庆-重庆",
            path: "#",
            describe: ""
        }, {
            name: "四川-成都",
            path: "#",
            describe: ""
        }]
    }, {
        name: "华东",
        path: "#",
        describe: "",
        children: [{
            name: "浙江-杭州",
            path: "#",
            describe: ""
        }, {
            name: "浙江-温州",
            path: "#",
            describe: ""
        }, {
            name: "江西-南昌",
            path: "#",
            describe: ""
        }, {
            name: "江西-赣州",
            path: "#",
            describe: ""
        }, {
            name: "上海-上海",
            path: "#",
            describe: ""
        }, {
            name: "江苏-南京",
            path: "#",
            describe: ""
        }, {
            name: "江苏-苏州",
            path: "#",
            describe: ""
        }]
    }, {
        name: "西北",
        path: "#",
        describe: "",
        children: [{
            name: "陕西-西安",
            path: "#",
            describe: ""
        }, {
            name: "新疆-哈密",
            path: "#",
            describe: ""
        }]
    }, {
        name: "东北",
        path: "#",
        describe: "",
        children: [{
            name: "辽宁-大连",
            path: "#",
            describe: ""
        }]
    }];

    useEffect(() => {
        Citizen.preproc(staticArches);
        Citizen.preproc(staticPivots);
        setArches(staticArches);
        setPivots(staticPivots);
    }, []);

    useEffect(() => {
        for (let arch of arches) {
            for (let subsystem of (arch.children || [])) {
                if (window.location.pathname && window.location.pathname.indexOf(Citizen.abs(subsystem)) > -1) {
                    setSubsystem(subsystem);
                    return;
                }
            }
        }
        setSubsystem(staticArches?.[0].children?.[0] || new Citizen());
    }, [arches]);

    const onClickPivot = (pivot: Citizen) => {

    }

    const onClickMenu = (item: Citizen) => {
        if (item?.parent?.parent) {
            setSubsystem(item?.parent?.parent);
        }
        setCrumbs(Citizen.chain(item));
    }

    const renderRoutes = (routes: Citizen[]) => {
        if (!routes || routes.length < 1) {
            return [];
        }
        return (
            <Routes>
                {
                    routes.filter(route => route.path.endsWith('/*')).map(route => {
                        return <Route key={route.path} path={route.path.replaceAll('/*', '')} element={route.route}/>
                    })
                }
                {
                    routes.map(route => {
                        return (
                            <Route key={route.path} path={route.path} element={
                                (route?.children?.length || 0) < 1 ? route.route : renderRoutes(route.children || [])
                            }/>
                        )
                    })
                }
            </Routes>
        )
    }

    return (
        <div>
            <Container root={"/mesh"} arches={arches} pivots={pivots} onClickPivot={onClickPivot}
                       onClickMenu={onClickMenu}>
                <Flex style={{height: '100%', width: '100%'}} vertical={false} justify={'start'}>
                    <SideMenu root={"/mesh"} subsystem={subsystem} onClick={onClickMenu}/>
                    <div/>
                    <div style={{width: '100%', height: '100%', padding: '0 1rem'}}>
                        <Flex vertical={false} align={'center'} className={styles.crumbView}>
                            <Breadcrumb items={
                                crumbs.map((crumb, idx) => {
                                    return {
                                        key: `${idx}`,
                                        title: crumb.name,
                                    }
                                })}/>
                            {
                                crumbs.filter((_, idx) => idx == crumbs.length - 1).map((crumb, idx) => (
                                    <Tooltip key={`${idx}`} placement="bottom" overlay={crumb.describe || ''}>
                                        <div style={{cursor: 'pointer', color: 'blue'}}>
                                            &nbsp;&nbsp;&nbsp;<QuestionCircleOutlined/>
                                        </div>
                                    </Tooltip>
                                ))
                            }
                        </Flex>
                        <Watermark style={{width: '100%', height: '100%'}} content={environ.name || ''}
                                   gap={[200, 200]}>
                            {renderRoutes(arches)}
                        </Watermark>
                    </div>
                </Flex>
            </Container>
        </div>
    )
}


