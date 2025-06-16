import React, { useEffect, useState } from "react";
import styles from "./index.module.less";
import BgSVG from '../../assets/bg.png';

const ExtractTaskPage: React.FC = () => {

    return (
        <div className={styles.container}
            style={{
                backgroundImage: `url(${BgSVG})`,
            }}
        >

            <div className={styles.header}>
            </div>
            <div className={styles.body}>
            </div>

            <div className={styles.footer}>
            </div>
        </div>
    )
};

export default ExtractTaskPage;
