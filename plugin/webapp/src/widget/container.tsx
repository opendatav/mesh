/*
 * Copyright (c) 2019, 2024, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

import {Avatar, Badge, Button, Divider, Drawer, Flex, Popover, Segmented, Tag} from "antd";
import React, {ReactNode, useEffect, useState} from "react";
import Icon, {BellOutlined, ContactsOutlined, GlobalOutlined, QuestionCircleOutlined} from "@ant-design/icons";
import {Citizen, Environ, Session} from "@ducesoft/mesh";
import {DropdownMenu, GridView, MailBox, PopupMenu, useStyles} from "@/widget";
import {Link} from "react-router-dom";
import {MeshIcon} from "@/assets";


function InfoCard(props: { session: Session }) {

    const {styles, cx, theme} = useStyles();
    const [session, setSession] = useState<Session>(props.session);
    const [tab, setTab] = useState('基础信息');

    const onTabChange = (v: string) => {
        setTab(v);
    }

    useEffect(() => {
        setSession(props.session);
    }, [props.session])

    return (
        <div style={{width: '300px', backgroundColor: theme.x.buttonAltBG}}>
            <Flex className={styles.padding20} vertical={true} style={{fontSize: '13px'}}>
                {
                    [
                        {k: '用户编号', v: session.userid},
                        {k: '用户名称', v: session.username}
                    ].map(kv => {
                        return (
                            <div key={kv.k} style={{marginBottom: '5px'}}>
                                <span className={styles.nameText}>{kv.k}</span>
                                <span className={styles.logonIdText}>&nbsp;&nbsp;&nbsp;{kv.v}</span>
                            </div>
                        )
                    })
                }
                <div style={{marginBottom: '5px'}}>
                    {
                        session.attrs?.map((attr, idx) => <Tag key={`${idx}`} color="blue">{attr}</Tag>)
                    }
                </div>
            </Flex>
            <div style={{backgroundColor: theme.colorBgBase, padding: '1px 10px 0 0'}}>
                <Segmented defaultValue={'基础信息'} options={['基础信息', '权限属性', '安全设置']} block
                           onChange={onTabChange}/>
                <Flex className={styles.padding20} vertical={true} style={{fontSize: '13px'}}>
                    {
                        [
                            {k: '联系电话', v: session.mobile},
                            {k: '电子邮箱', v: session.mail},
                            {k: '登记枢纽', v: session.nodeId}
                        ].map(kv => {
                            return (
                                <div key={kv.k} style={{marginBottom: '5px'}}>
                                    <span className={styles.nameText}>{kv.k}</span>
                                    <span className={styles.logonIdText}>&nbsp;&nbsp;&nbsp;{kv.v}</span>
                                </div>
                            )
                        })
                    }
                </Flex>
            </div>
            <div className={styles.padding20} style={{backgroundColor: theme.colorBgBase}}>
                <Button block>退出登录</Button>
            </div>
        </div>
    )
}

export function Container(props: {
    root: string,
    pivots: Citizen[],
    arches: Citizen[],
    onClickPivot: (pivot: Citizen) => void,
    onClickMenu: (item: Citizen) => void,
    children: ReactNode,
}) {

    const {styles, cx, theme} = useStyles();
    const [environ, setEnviron] = useState(new Environ());
    const [openPopup, setOpenPopup] = useState(false);
    const [session, setSession] = useState<Session>(new Session());
    const [contacts, setContacts] = useState<Citizen[]>([]);
    const staticSession: Session = {
        attrs: ["系统管理员", "运维人员"],
        avatar: "https://oss.aliyuncs.com/aliyun_id_photo_bucket/default_handsome.jpg",
        deptId: "",
        raw: {},
        instIds: [],
        mail: "coyzeng@gmail.com",
        mobile: "13834567876",
        nickname: "CoyZeng",
        nodeId: "浙江-杭州",
        token: "",
        userid: "10000001",
        username: "CoyZeng"
    };
    const staticContacts: Citizen[] = [{
        name: "管理办公室",
        path: "#",
        describe: "",
        children: [{
            name: <span>0618-988677&nbsp;&nbsp;&nbsp;</span>,
            path: "#",
            describe: ""
        }, {
            name: <span>0618-988677&nbsp;&nbsp;&nbsp;</span>,
            path: "#",
            describe: ""
        }, {
            name: <span>0618-988677&nbsp;&nbsp;&nbsp;</span>,
            path: "#",
            describe: ""
        }]
    }, {
        name: "运维办公室",
        path: "#",
        describe: "",
        children: [{
            name: <span>0618-988677&nbsp;&nbsp;&nbsp;</span>,
            path: "#",
            describe: ""
        }, {
            name: <span>0618-988677&nbsp;&nbsp;&nbsp;</span>,
            path: "#",
            describe: ""
        }, {
            name: <span>0618-988677&nbsp;&nbsp;&nbsp;</span>,
            path: "#",
            describe: ""
        }]
    }];

    const onHoverPopupMenu = () => {
        setOpenPopup(!openPopup);
    }

    const onClickPivot = (pivot: Citizen) => {
        props?.onClickPivot?.(pivot);
    }

    const onClickMenu = (item: Citizen) => {
        props?.onClickMenu?.(item);
    }

    useEffect(() => {
        setSession(staticSession);
        setContacts(staticContacts);
    }, []);

    return (
        <div>
            <div className={styles.stickyView}>
                <Flex gap="middle" vertical={false} align={'center'} justify={'space-between'}>
                    <Flex vertical={false} align={'center'} justify={'center'}>
                        <button className={styles.menuButton}
                                onClick={onHoverPopupMenu}
                                onMouseEnter={onHoverPopupMenu}>
                            <Flex vertical={true} align={'center'}>
                                    <span className={styles.buttonIcon}
                                          style={openPopup ? {
                                              transform: 'rotate(45deg)',
                                              marginLeft: '4px'
                                          } : {transform: 'rotate(0deg)'}}/>
                                <span className={styles.buttonIcon} style={{opacity: openPopup ? 0 : 1}}/>
                                <span className={styles.buttonIcon}
                                      style={openPopup ? {
                                          transform: 'rotate(-45deg)',
                                          marginLeft: '4px'
                                      } : {transform: 'rotate(0deg)'}}/>
                            </Flex>
                        </button>
                        <Icon component={MeshIcon} style={{fontSize: theme.x.appbarIconSize, marginRight: '4px'}}/>
                        <div className="text-xl font-bold">{environ.name || '服务网格平台'}</div>
                        <a href="https://aaas.firmer.tech" target="_blank" className={styles.linkButton}>
                            <GlobalOutlined style={{marginRight: '8px'}}/>
                            <span style={{fontSize: '12px', lineHeight: '30px'}}>应用服务平台</span>
                        </a>
                        <DropdownMenu items={props?.pivots || []} onClick={onClickPivot}/>
                    </Flex>
                    <Flex vertical={false} align={'center'}>
                        <Popover placement="bottomRight" content={<GridView items={contacts}/>} arrow={false}
                                 styles={{body: {borderRadius: '0'}}}>
                            <div className={styles.iconButton}>
                                <ContactsOutlined/>
                            </div>
                        </Popover>
                        <Popover placement="bottom" content={<MailBox/>} arrow={false}
                                 styles={{body: theme.x.popoverBorderStyle}}>
                            <Badge size={'small'} count={5} overflowCount={99} color={'red'}
                                   className={styles.iconButton}>
                                <BellOutlined/>
                            </Badge>
                        </Popover>
                        <Popover placement="bottom"
                                 content={<Link to={'#'} className={styles.linkText}>使用文档</Link>}
                                 arrow={false}
                                 styles={{body: {borderRadius: '0'}}}>
                            <div className={styles.iconButton}>
                                <QuestionCircleOutlined/>
                            </div>
                        </Popover>
                        <Divider style={{height: '100%'}} type={'vertical'}/>
                        <Popover placement="bottomRight" content={<InfoCard session={session}/>} arrow={false}
                                 styles={{body: theme.x.popoverBorderStyle}}>
                            <Flex vertical={false} align={'center'} justify={'space-between'}
                                  style={{padding: '0 15px 0 0', cursor: 'pointer'}}>
                                <Flex vertical={true} align={'flex-end'} style={{margin: '0 10px 0 0'}}>
                                    <div className={styles.logonIdText}>{session.userid}</div>
                                    <div className={styles.nameText}>{session.nickname}</div>
                                </Flex>
                                <Avatar shape={'circle'} src={"" == session.avatar ? null : session.avatar}/>
                            </Flex>
                        </Popover>
                    </Flex>
                </Flex>
            </div>
            <div className={styles.drawerView}>
                <div className={styles.drawerView}>
                    {props?.children}
                </div>
                <Drawer
                    styles={{body: {padding: 0}}}
                    width={'45vw'}
                    placement="left"
                    closable={false}
                    maskClosable={true}
                    open={openPopup}
                    onClose={onHoverPopupMenu}
                    getContainer={false}>
                    <div><PopupMenu className={styles.drawerView}
                                    root={props?.root}
                                    items={props?.arches || []}
                                    onClick={onClickMenu}
                                    onClose={onHoverPopupMenu}/>
                    </div>
                </Drawer>
            </div>
        </div>
    )
}
