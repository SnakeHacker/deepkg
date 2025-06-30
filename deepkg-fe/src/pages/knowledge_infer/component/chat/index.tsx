import React, { useEffect } from "react";
import styles from "./index.module.less";
import { Welcome, Bubble, Sender, useXAgent, useXChat } from "@ant-design/x";
import { Button, Flex, type GetProp, Typography } from 'antd';
import type { BubbleProps } from '@ant-design/x';
import markdownit from 'markdown-it';
import { UserOutlined } from '@ant-design/icons';
import { baseUrl } from "../../../../utils/req";
import customIcon from '../../../../assets/monkey.png';

const md = markdownit({ html: true, breaks: true });

const roles: GetProp<typeof Bubble.List, 'roles'> = {
    ai: {
      placement: 'start',
      avatar: { icon: <img src={customIcon}/> , style: { background: '#ebf7ff' } },

    },
    local: {
      placement: 'end',
      avatar: { icon: <UserOutlined />, style: { background: '#87d068' } },
    },
  };

const renderMarkdown: BubbleProps['messageRender'] = (content: any) => {
    return (
        <Typography>
            <div dangerouslySetInnerHTML={{ __html: md.render(content) }} />
        </Typography>
    );
};


interface ChatProps {
}

const ChatContainer: React.FC<ChatProps> = ({
  }) =>  {

  useEffect(() => {

  }, []);

  const [content, setContent] = React.useState('');

  // Agent for request
  const [agent] = useXAgent({
    request: async ({ messages, message }, { onSuccess, onUpdate }) => {

      const history: any[] = messages!.map((content, index) => {
        console.log(content)
        return {
            role: index % 2 === 0 ? 'user' : 'assistant',
            content: content
        };
      });

      history.pop()

      var raw = JSON.stringify({
        "content": message,
        "history": history,
      });

      var myHeaders = new Headers();
      myHeaders.append("Content-Type", "application/json");
      myHeaders.append("Accept", "text/event-stream");

      const response = await fetch(baseUrl+"/api/chat",
        {
            method: 'POST',
            headers: myHeaders,
            body: raw,
        }
      )

      var buffer = '';
      const reader = response.body!.pipeThrough(new TextDecoderStream()).getReader();

      function streamingRead(inReader: any, buf: string ){
        inReader.read().then(({done, value}: any)=>{
            if (done){
                console.log('Stream complete');
                onSuccess(buf as any);
                buf = '';
                return
            }

            const lines = value.trim().split('\n');
            lines.forEach((line: string) => {

                // 使用正则表达式匹配 JSON 格式的内容
                const match = line.match(/data: (\{.*\})/);
                if (match) {
                  // 提取 JSON 字符串并解析为对象
                  const jsonData = JSON.parse(match[1]);
                  var resultValue = jsonData.result;
                  if (resultValue.endsWith("[DONE]")){
                    return
                  }


                  if (resultValue == '<think>') {
                    buf += "<b>思维链: </b> \n"
                    setContent(buf);
                    return
                  }

                  if (resultValue == '</think>'){

                    buf += "\n\r\n\r\n\r\n\r___"
                    setContent(buf);
                    return
                 }

                  buf += resultValue

                  onUpdate(buf as any);
                }
              });

            return streamingRead(inReader, buf);
          })
      }

      streamingRead(reader, buffer)
    },
  });

  const clearMessages = () => {
    setMessages([]);
  };

  const { onRequest, messages, setMessages } = useXChat({
    agent,
  });


  return (
    <div className={styles.chatContainer}>
        <Flex vertical gap="middle">
            <div className={ styles.header}>
                <Welcome
                    style={{
                        background: 'linear-gradient(97deg, rgba(90,196,255,0.12) 0%, rgba(174,136,255,0.12) 100%)',
                        borderStartStartRadius: 4,
                    }}
                    icon={<img src={customIcon} style={{
                        width: "80px",
                        height: "80px",
                        borderRadius: "50%" // 设置为圆形
                    }}/>}
                    title="您好！我是“KG助手”"
                    description="在这里，我将用专业的知识，为您提供贴心的服务。"
                />

                <Button onClick={clearMessages} type="primary" className={styles.newChat}>
                    新对话
                </Button>
            </div>

            <Bubble.List
                roles={roles}
                style={{ maxHeight: '60vh', minHeight: '60vh' }}
                items={messages.map(({ id, message, status }) => ({
                    key: id,
                    role: status === 'local' ? 'local' : 'ai',
                    // content:renderMarkdown(message),
                    content: message,
                }))}
            />
            <Sender
                loading={agent.isRequesting()}
                value={content}
                onChange={setContent}
                onSubmit={(nextContent) => {
                    onRequest(nextContent);
                    setContent('');
                }}
            />
        </Flex>
    </div>
  )
};

export default ChatContainer;
