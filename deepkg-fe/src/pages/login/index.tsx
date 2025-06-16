import React, { useEffect, useState } from 'react';
import { Form, Input, Button, message } from 'antd';
import styles from './index.module.less';
import { isAuthenticated, setAuthenticated } from '../../service/auth';
import { GetCaptcha, GetPublicKey, Login } from '../../service/session';
import {JSEncrypt} from 'jsencrypt';
import { useNavigate } from 'react-router-dom';
// import LogoSVG from '../../assets/logo.svg';
import BgSVG from '../../assets/background.png';


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

        console.log(res)

        if (res.id != null) {
            console.log(res)
            messageApi.success('登录成功');
            setAuthenticated(res)
            navigate('/')
        } else {
            messageApi.error(res.returnDesc);
            getCaptcha();
        }
    };

    const [init, setInit] = useState(false);


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
            <img src={BgSVG} className={styles.backgroundImage}/>

            {/* <div style={{ width: '100%', height: '50px', top: '10%', position: 'absolute' }}>
                <Orb
                    hoverIntensity={0.5}
                    rotateOnHover={true}
                    hue={0}
                    forceHoverState={false}
                />
            </div> */}
            <div className={styles.logoContainer}>
                {/* <img src={LogoSVG} className={styles.logo} /> */}
                <h1 className={styles.loginTitle}>知识图谱平台</h1>

                <Form
                    className={styles.loginForm}
                    name="login"
                    initialValues={{ remember: true }}
                    onFinish={onFinish}
                    autoComplete="off"
                >
                    <div style={{marginBottom:'10px'}}>
                        用户名
                    </div>
                    <Form.Item
                        name="account"
                        rules={[{ required: true, message: '请输入用户名!' }]}
                    >
                        <Input
                            style={{height: '48px'}}
                            placeholder="用户名"
                        />
                    </Form.Item>
                    <div style={{marginBottom:'10px'}}>
                        密码
                    </div>
                    <Form.Item
                        name="password"
                        rules={[{ required: true, message: '请输入密码!' }]}
                    >
                        <Input.Password
                            style={{height: '48px'}}
                            placeholder="密码"
                        />
                    </Form.Item>
                    <div style={{marginBottom:'10px'}}>
                        验证码
                    </div>
                    <Form.Item
                        name="captcha_value"
                        rules={[{ required: true, message: '请输入验证码!' }]}
                        style={{ display: 'flex', alignItems: 'center' }}
                    >
                        <div className={styles.captcha}>
                            <Input
                                style={{ height:'48px', flex: 1, marginRight: 10 }}
                                placeholder="验证码"
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
                            style={{background: '#1F35DB', color: 'white', height: '48px'}}
                            variant="solid"
                            htmlType="submit"
                            className={styles.loginFormButton}
                        >
                            登录
                        </Button>
                    </Form.Item>
                </Form>
            </div>
        </div>
    );
};

export default LoginPage;