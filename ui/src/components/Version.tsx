import { memo, useEffect, useState } from "react";
import { useTranslation } from "react-i18next";
import { ReadOutlined as ReadOutlinedIcon } from "@ant-design/icons";
import { Badge, Divider, Space, Typography } from "antd";

import { version } from "@/domain/version";
import { useVersionChecker } from "@/hooks";

export type VersionProps = {
  className?: string;
  style?: React.CSSProperties;
};

const Version = ({ className, style }: VersionProps) => {
  const { t } = useTranslation();

  const { data: hasNewVersions } = useVersionChecker();

  return (
    <Space className={className} style={style} size={4}>
      <Typography.Link type="secondary" href="https://docs.certimate.me" target="_blank">
        <div className="flex items-center justify-center space-x-1">
          <ReadOutlinedIcon />
          <span>{t("common.menu.document")}</span>
        </div>
      </Typography.Link>

      <Divider type="vertical" />

      <Badge styles={{ indicator: { transform: "scale(0.75) translate(50%, -50%)" } }} count={hasNewVersions ? "NEW" : 0}>
        <Typography.Link type="secondary" href="https://github.com/usual2970/certimate/releases" target="_blank">
          {version}
        </Typography.Link>
      </Badge>
    </Space>
  );
};

export default memo(Version);
