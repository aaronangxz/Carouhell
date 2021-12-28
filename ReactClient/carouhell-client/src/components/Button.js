import React from 'react';
import { Button } from 'tdesign-react';

const ButtonExample = () => {
  return (
    <div className="tdesign-demo-block-column">
      <div className="tdesign-demo-block-row">
        <Button theme="default" variant="base">填充按钮</Button>
        <Button theme="default" variant="outline">描边按钮</Button>
        <Button theme="default" variant="dashed">虚框按钮</Button>
        <Button theme="default" variant="text">文字按钮</Button>
        <Button loading theme="default">加载中</Button>
      </div>
    </div>
  )
}

export default ButtonExample;
