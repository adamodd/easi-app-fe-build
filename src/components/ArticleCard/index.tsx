import React from 'react';
import { useTranslation } from 'react-i18next';
import { useHistory } from 'react-router-dom';
import {
  Card,
  CardBody,
  CardFooter,
  CardHeader
} from '@trussworks/react-uswds';
import classnames from 'classnames';

import UswdsReactLink from 'components/LinkWrapper';
import Tag from 'components/shared/Tag';
import { ArticleProps } from 'types/articles';

import './index.scss';

type ArticleCardProps = {
  className?: string;
  isLink?: boolean;
};

const ArticleCard = ({
  className,
  type,
  route,
  translation,
  isLink = false
}: ArticleCardProps & ArticleProps) => {
  const { t } = useTranslation(translation);
  const history = useHistory();

  const clickHandler = (url: string) => {
    if (isLink) {
      history.push(url);
    }
  };

  return (
    <Card
      containerProps={{ className: 'radius-md shadow-2' }}
      data-testid="article-card"
      className={classnames('desktop:grid-col-4', className, {
        'article-card--isLink': isLink
      })}
      onClick={() => clickHandler(`help${route}`)}
    >
      <CardHeader className="padding-x-3 padding-top-3 padding-bottom-2">
        <h3 className="line-height-body-4 margin-bottom-1">{t('title')}</h3>
        <Tag className="system-profile__tag text-primary-dark bg-primary-lighter padding-bottom-1">
          {type}
        </Tag>
      </CardHeader>
      <CardBody className="padding-y-0 article__body">
        {t('description')}
      </CardBody>
      <CardFooter className="padding-top-2">
        <UswdsReactLink
          to={`help${route}`}
          className="easi-request__button-link padding-x-2"
        >
          {useTranslation('help').t('read')}
        </UswdsReactLink>
      </CardFooter>
    </Card>
  );
};

export default ArticleCard;