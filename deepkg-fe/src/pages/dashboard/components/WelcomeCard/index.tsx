import React from 'react';
import { Card, Button, Divider, Typography, Avatar } from 'antd';
import { useNavigate } from 'react-router-dom';
import { getCurrentUser } from '../../../../service/auth';
import styles from './index.module.less';

const { Title, Text } = Typography;

const WelcomeCard: React.FC = () => {
    const user = getCurrentUser();
    const navigate = useNavigate();

    const getTimeGreeting = () => {
        const hour = new Date().getHours();
        if (hour < 12) return "上午好";
        if (hour < 18) return "下午好";
        return "晚上好";
    };
    const motivationalQuotes = [
        "风起时学会顺风而行，雨落时记得为自己撑起一方晴空。",
        "别被琐碎日常耗尽热爱，人生值得你满怀期待地奔赴每一天。",
        "照顾好身体，也别忘了家人和自己，生活才会回馈你真实的温柔。",
        "愿你生活平凡但安稳，收入稳定，身边人常伴。",
        "努力让自己变强，是对未来最温柔的负责。",
        "幸福从不遥远，它藏在眼里的风景、碗里的饭菜与身边的人。",
        "清晨有希望，傍晚有归期，日子平淡也温馨。",
        "别慌，人生不会一直晴朗，但总有月光为你照亮路途。",
        "走过风雪，愿你回头时依旧温柔笃定，不负自己。",
    ];

    const getRandomQuote = () => {
        const index = Math.floor(Math.random() * motivationalQuotes.length);
        return motivationalQuotes[index];
    };

    const [quote] = React.useState(getRandomQuote());

    return (
        <Card className={styles.welcomeCard}>
            <div className={styles.container}>
                <Avatar
                    size={64}
                    src={user?.avatar}
                    className={styles.avatar}
                    icon={!user?.avatar && user?.username?.[0]}
                />
                <div className={styles.textContent}>
                    <Title level={4} className={styles.greeting}>
                        {getTimeGreeting()}，{user?.nickName || user?.username || '用户'}，祝你开心每一天！
                    </Title>
                    <Text type="secondary">
                        <span style={{ color: '#3d5bf1' }}>系统管理员</span>
                        <Divider type="vertical" />
                        {quote}
                    </Text>
                </div>
                <div className={styles.action}>
                    <Button type="primary" ghost size="large" onClick={() => navigate('/user')}>
                        个人中心
                    </Button>
                </div>
            </div>
        </Card>
    );
};

export default WelcomeCard;
