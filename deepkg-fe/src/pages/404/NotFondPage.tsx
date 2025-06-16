import React from "react";
import notFoundSVG from "../../assets/404.svg";

const NotFoundPage: React.FC = () => {
  return (
    <div>
      <img src={notFoundSVG} style={{ height: '99vh' }} />
    </div>
  );

};

export default NotFoundPage;