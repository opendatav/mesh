import type {ReactNode} from 'react';
import clsx from 'clsx';
import Heading from '@theme/Heading';
import styles from './styles.module.css';

type FeatureItem = {
    title: string;
    Svg: React.ComponentType<React.ComponentProps<'svg'>>;
    description: ReactNode;
};

const FeatureList: FeatureItem[] = [
    {
        title: '隐私计算全栈互联',
        Svg: require('@site/static/img/undraw_docusaurus_mountain.svg').default,
        description: (
            <>
                构建基于联邦学习、多方安全计算（MPC）、可信执行环境（TEE）等隐私计算技术的统一接入和协同网络接入平台。
                提供标准化接口与协议，实现不同隐私计算平台之间的互通互联。
                打破不同隐私计算技术形成的新孤岛，实现数据“可用不可见”，打破数据孤岛。
            </>
        ),
    },
    {
        title: '弹性可扩展的金融级通信代理',
        Svg: require('@site/static/img/undraw_docusaurus_tree.svg').default,
        description: (
            <>
                基于服务网格（Service Mesh）和云原生架构构建高可用、低延迟的通信代理层。
                支持动态插拔各类中间件和服务插件（如加密、压缩、流量控制等），适应不同业务需求。
                多活部署、故障隔离、自动熔断机制保障系统稳定运行。
                保障金融机构间异构系统的安全高效对接，快速集成新业务通道或第三方服务商。
            </>
        ),
    },
    {
        title: '全链路服务治理可观测',
        Svg: require('@site/static/img/undraw_docusaurus_react.svg').default,
        description: (
            <>
                基于服务网格技术实现服务级流量管理、策略控制与安全保障。
                全链路追踪、指标监控、日志聚合实现Tracing/Metrics/Logging三位一体，精细化运维与问题定位。
                支持多租户隔离、权限控制、API网关、限流降级等功能。
                保障金融交易、支付清算等关键业务的全流程可视化与异常检测，提升系统透明度与风险可控性。
            </>
        ),
    },
];

function Feature({title, Svg, description}: FeatureItem) {
    return (
        <div className={clsx('col col--4')}>
            <div className="text--center">
                <Svg className={styles.featureSvg} role="img"/>
            </div>
            <div className="text--center padding-horiz--md">
                <Heading as="h3">{title}</Heading>
                <p>{description}</p>
            </div>
        </div>
    );
}

export default function HomepageFeatures(): ReactNode {
    return (
        <section className={styles.features}>
            <div className="container">
                <div className="row">
                    {FeatureList.map((props, idx) => (
                        <Feature key={idx} {...props} />
                    ))}
                </div>
            </div>
        </section>
    );
}
