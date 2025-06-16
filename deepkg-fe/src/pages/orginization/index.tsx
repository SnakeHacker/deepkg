import React, { useEffect, useState } from 'react';
import { PlusOutlined, RollbackOutlined } from '@ant-design/icons';
import { CreateOrg, DeleteOrgs, ListOrg, UpdateOrg } from '../../service/organization';
import { Button, Input, Modal, Pagination, Popconfirm, Table } from 'antd';
import styles from './index.module.less';
import BgSVG from '../../assets/bg.png';
import type { Organization } from '../../model/organization';
const OrganizationListPage: React.FC = () => {
    const [pagination, setPagination] = useState({
          current: 1,
          pageSize: 10,
    });
    const [total, setTotal] = useState(0);
    const [orgs, setOrgs] = useState<Organization[]>([]);
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [orgName, setOrgName] = useState('');
    const [orgID, setOrgID] = useState(0);

    useEffect( () => {
      listOrgs();
    }, [pagination]);

    const listOrgs = async () => {
      const res = await ListOrg({
          page_size: pagination.pageSize,
          page_number: pagination.current
      });
      console.log(res)

      setTotal(res.total)

      setOrgs(res.organizations)

    };

    const handlePageChange = (page: number, pageSize?: number) => {
      setPagination(prev => ({
        ...prev,
        current: page,
        pageSize: pageSize || prev.pageSize
      }));
    };

    const handleCreateOrgOk = () => {
        setIsModalOpen(false);
        if (orgID > 0){
            updateOrg().then((res)=>{
                console.log(res)
                setOrgName('')
                setOrgID(0)
                listOrgs()
            })
        }else{
            createOrg().then((res)=>{
                console.log(res)
                setOrgName('')
                setOrgID(0)
                listOrgs()
            })
        }

    }

    const handleCancelCreateOrg = () => {
        setIsModalOpen(false);
        setOrgName('')
        setOrgID(0)
    }

    const createOrg = async ()=>{
        const res = await CreateOrg({
            "organization":{
                "org_name": orgName,
            }
        })
        return res
    }

    const updateOrg = async ()=>{
        const res = await UpdateOrg({
            "organization":{
                "id": orgID,
                "org_name": orgName,
            }
        })
        return res
    }
    const deleteOrg =async (id: number )=>{
        const res = await DeleteOrgs({ ids: [id] });
        console.log(res)
        listOrgs();
    }

    const columns = [
      {
        title: 'ID',
        dataIndex: 'id',
        key: 'id',
        width: 50,
      },
      {
        title: '组织名称',
        dataIndex: 'org_name',
        key: 'org_name',
        width: 350,
      },
      {
        title: '创建时间',
        dataIndex: 'created_at',
        key: 'created_at',
        width: 350,
      },
      {
        title: '更新时间',
        dataIndex: 'updated_at',
        key: 'updated_at',
        width: 350,
      },
      {
        title: '操作',
        key: 'action',
        width: 200,
        render: (_: any, record:Organization) => (
          <div>
              <Button
                  style={{ marginRight: '10px' }}
                  onClick={()=>{
                    setOrgID(record.id!)
                    setOrgName(record.org_name)
                    setIsModalOpen(true)
                  }}
                  size='small'
              >
                编辑
              </Button>

              <Popconfirm
                title={""}
                description="确认是否删除组织?"
                onConfirm={()=>deleteOrg(record.id!)}
                onCancel={()=>{}}
                okText="确认"
                cancelText="取消"
              >
                <Button
                    size='small'
                    danger
                >
                    删除
                </Button>
              </Popconfirm>

          </div>
        ),
      },
    ];

    return (
      <div className={styles.container}
        style={{
            backgroundImage: `url(${BgSVG})`,
        }}
      >

        <div className={styles.header}>
          <Button
            style={{ marginRight: '10px' }}
            type="default"
            icon={<RollbackOutlined />}
            onClick={() => {
                window.history.back();
            }}
            >
            返回
          </Button>
          <Button
            type="primary"
            icon={<PlusOutlined />}
            onClick={
                () => {setIsModalOpen(true)}}
          >
            新建组织
          </Button>
        </div>
        <div className={styles.body}>
            <Modal
                title={
                    orgID > 0 ? '编辑组织' : `新建组织`
                }
                open={isModalOpen}
                onOk={handleCreateOrgOk}
                onCancel={handleCancelCreateOrg}
                okText="确定"
                cancelText="取消"
                okButtonProps={{
                    disabled: orgName.length == 0
                }}
                width={350}
            >
                <div className={styles.modal_container}>
                    <Input
                        placeholder="请输入组织名称"
                        value={orgName}
                        style={{'width':'300px'}}
                        onChange={(e)=>setOrgName(e.target.value)}
                    />
                </div>
            </Modal>

            <Table
                style={{ width: '100%' }}
                dataSource={orgs}
                columns={columns}
                pagination={false}
                rowKey="id"
            />
        </div>

        <div className={styles.footer}>
          <Pagination
            current={pagination.current}
            pageSize={pagination.pageSize}
            total={total}
            onChange={handlePageChange}
          />
        </div>
      </div>
    );
  };


export default OrganizationListPage;