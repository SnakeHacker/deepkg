import React, {useEffect, useState} from 'react';
import {Button, Col, Form, Input, message, Row, Tabs} from 'antd';
import styles from './index.module.less';
import {isAuthenticated, setAuthenticated} from '../../service/auth';
import {GetCaptcha, GetPublicKey, Login} from '../../service/session';
import {JSEncrypt} from 'jsencrypt';
import {useNavigate} from 'react-router-dom';
import LogoHorizontal from '../../assets/logo_horizontal.png';
import LeftBg from '../../assets/left_bg.png';
import {KeyOutlined, SafetyOutlined, UserOutlined} from "@ant-design/icons";


const LoginPage: React.FC = () => {
    const [messageApi, contextHolder] = message.useMessage();
    let navigate = useNavigate();
    const [captchaID, setCaptchaId] = useState('');
    const [captchaBase64, setCaptchaBase64] = useState('');

    const onFinish = (values: { account: string; password: string, captcha_value: string }) => {
        startLogin(values.account, values.password, values.captcha_value);
    };

    const startLogin =async  (account: string, password: string, captcha_value: string) => {

        const publicKeyRes = await GetPublicKey();
        const encryptor = new JSEncrypt();
        encryptor.setPublicKey(publicKeyRes.public_key);
        const encryptedPwd = encryptor.encrypt(password.trim());

        const payload = {
            account: account,
            password: encryptedPwd.toString(),
            captcha_id: captchaID,
            captcha_value: captcha_value,
        }

        const res = await Login(payload)

        if (res.id != null) {
            messageApi.success('登录成功');
            setAuthenticated(res)
            navigate('/')
        } else {
            messageApi.error(res.returnDesc);
            getCaptcha();
        }
    };

    useEffect(() => {
        console.log('%c:)', 'color: green; font-size: 32px;');

        if (isAuthenticated()) {
            navigate('/')
        }

    }, []);

    useEffect(() => {
        getCaptcha();
    }, [])

    const getCaptcha = async () => {
        const res = await GetCaptcha();
        setCaptchaId(res.captcha_id)
        setCaptchaBase64(res.captcha_base64)

    };

    return (
        <div className={styles.container}>
            {contextHolder}
            <Row align="middle" style={{ height: '100%' }}>
                <Col span={12} className={styles.leftSide}>
                    <img src={LeftBg} alt="background" className={styles.leftImage} />
                </Col>

                {/* 右侧登录表单 */}
                <Col span={12} className={styles.rightSide}>
                    <img src={LogoHorizontal} alt="background" className={styles.logo} />
                    <div className={styles.logoContainer}>

                        <Tabs centered size={"large"} items={[
                            {
                                key: "1",
                                label: "账号登录",
                                children: (<Form
                                    className={styles.loginForm}
                                    name="login"
                                    initialValues={{ remember: true }}
                                    onFinish={onFinish}
                                    autoComplete="off"
                                >
                                    <Form.Item
                                        name="account"
                                        rules={[{ required: true, message: '请输入用户名!' }]}
                                    >
                                        <Input
                                            style={{height: '48px'}}
                                            placeholder="用户名"
                                            prefix={<UserOutlined/>}
                                        />
                                    </Form.Item>
                                    <Form.Item
                                        name="password"
                                        rules={[{ required: true, message: '请输入密码!' }]}
                                    >
                                        <Input.Password
                                            style={{height: '48px'}}
                                            placeholder="密码"
                                            prefix={<KeyOutlined/>}
                                        />
                                    </Form.Item>
                                    <Form.Item
                                        name="captcha_value"
                                        rules={[{ required: true, message: '请输入验证码!' }]}
                                        style={{ display: 'flex', alignItems: 'center' }}
                                    >
                                        <div className={styles.captcha}>
                                            <Input
                                                style={{ height:'48px', flex: 1, marginRight: 10 }}
                                                placeholder="验证码"
                                                prefix={<SafetyOutlined/>}
                                            />
                                            {captchaBase64 && <img
                                                onClick={getCaptcha}
                                                src={captchaBase64}
                                                style={{ height: 45 }}
                                            />}
                                        </div>
                                    </Form.Item>
                                    <Form.Item>
                                        <Button
                                            style={{background: '#408aff', color: 'white', height: '48px'}}
                                            variant="solid"
                                            htmlType="submit"
                                            className={styles.loginFormButton}
                                        >
                                            登录
                                        </Button>
                                    </Form.Item>
                                </Form>)
                            }
                        ]}>
                        </Tabs>

                    </div>
                </Col>
            </Row>
        </div>
    );
};

export default LoginPage;