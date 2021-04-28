import React from 'react';
import { useTranslation } from 'react-i18next';
import { useHistory, useParams } from 'react-router-dom';
import { useMutation, useQuery } from '@apollo/client';
import { Alert } from '@trussworks/react-uswds';
import { DateTime } from 'luxon';
import GetAccessibilityRequestQuery from 'queries/GetAccessibilityRequestQuery';
import {
  GetAccessibilityRequest,
  GetAccessibilityRequest_accessibilityRequest_testDates as TestDateType
} from 'queries/types/GetAccessibilityRequest';
import { UpdateTestDate } from 'queries/types/UpdateTestDate';
import UpdateTestDateQuery from 'queries/UpdateTestDateQuery';

import useConfirmationText from 'hooks/useConfirmationText';
import { TestDateForm } from 'types/accessibilityRequest';
import { formatDate } from 'utils/date';

import Form from './Form';

import './styles.scss';

const TestDate = () => {
  const { t } = useTranslation('accessibility');
  const { confirmationText } = useConfirmationText();
  const { accessibilityRequestId, testDateId } = useParams<{
    accessibilityRequestId: string;
    testDateId: string;
  }>();
  const [mutate, mutateResult] = useMutation<UpdateTestDate>(
    UpdateTestDateQuery,
    {
      errorPolicy: 'all'
    }
  );
  const { data, loading } = useQuery<GetAccessibilityRequest>(
    GetAccessibilityRequestQuery,
    {
      variables: {
        id: accessibilityRequestId
      }
    }
  );
  const history = useHistory();

  const test: TestDateType = data?.accessibilityRequest?.testDates.find(
    date => date.id === testDateId
  );
  const testDate = DateTime.fromISO(test?.date);
  const initialValues: TestDateForm = {
    dateMonth: String(testDate.month),
    dateDay: String(testDate.day),
    dateYear: String(testDate.year),
    testType: test?.testType,
    score: {
      isPresent: Boolean(test?.score),
      value: test?.score ? `${test?.score && test?.score / 10}` : ''
    }
  };

  const onSubmit = (values: TestDateForm) => {
    const date = DateTime.fromObject({
      day: Number(values.dateDay),
      month: Number(values.dateMonth),
      year: Number(values.dateYear)
    });
    const hasScore = values.score.isPresent;
    const score = values.score.value;

    const confirmation = `
      ${t('testDateForm.confirmation.date', { date: formatDate(date) })}
      ${hasScore ? t('testDateForm.confirmation.score', { score }) : ''}
      ${t('testDateForm.confirmation.update')}
    `;

    mutate({
      variables: {
        input: {
          date,
          score: hasScore ? Math.round(parseFloat(score) * 10) : null,
          testType: values.testType,
          id: testDateId
        }
      }
    }).then(() => {
      history.push(`/508/requests/${accessibilityRequestId}`, {
        confirmationText: confirmation
      });
    });
  };

  if (loading) {
    return <p>Loading...</p>;
  }

  return (
    <>
      {confirmationText && (
        <Alert className="margin-top-4" type="success" role="alert">
          {confirmationText}
        </Alert>
      )}
      <Form
        initialValues={initialValues}
        onSubmit={onSubmit}
        error={mutateResult.error}
        requestName={data?.accessibilityRequest?.name}
        requestId={accessibilityRequestId}
        formType="update"
      />
    </>
  );
};

export default TestDate;
