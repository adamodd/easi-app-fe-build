import React from 'react';
import { useTranslation } from 'react-i18next';
import { useDispatch } from 'react-redux';
import { useHistory, useParams } from 'react-router-dom';
import { Button } from '@trussworks/react-uswds';
import { Field, Form, Formik, FormikProps } from 'formik';

import { ErrorAlert, ErrorAlertMessage } from 'components/shared/ErrorAlert';
import FieldErrorMsg from 'components/shared/FieldErrorMsg';
import FieldGroup from 'components/shared/FieldGroup';
import HelpText from 'components/shared/HelpText';
import Label from 'components/shared/Label';
import TextField from 'components/shared/TextField';
import { saveSystemIntake } from 'types/routines';
import { SubmitDatesForm, SystemIntakeForm } from 'types/systemIntake';
import flattenErrors from 'utils/flattenErrors';

const Dates = ({ systemIntake }: { systemIntake: SystemIntakeForm }) => {
  const { systemId } = useParams();
  const dispatch = useDispatch();
  const history = useHistory();
  const { t } = useTranslation();

  // TODO: Fix Text Field so we don't have to set initial empty values
  const initialValues: SubmitDatesForm = {
    grtDateDay: '',
    grtDateMonth: '',
    grtDateYear: '',
    grbDateDay: '',
    grbDateMonth: '',
    grbDateYear: ''
  };

  const onSubmit = (values: SubmitDatesForm) => {
    const {
      grtDateDay,
      grtDateMonth,
      grtDateYear,
      grbDateMonth,
      grbDateDay,
      grbDateYear
    } = values;
    const grtDate = `${grtDateYear}-${grtDateMonth}-${grtDateDay}`;
    const grbDate = `${grbDateYear}-${grbDateMonth}-${grbDateDay}`;

    const data = { ...systemIntake, grtDate, grbDate, id: systemId };
    console.log(data);
    dispatch(saveSystemIntake({ ...data }));

    history.push(`/governance-review-team/${systemId}/intake-request`);
  };

  return (
    <Formik
      initialValues={initialValues}
      onSubmit={onSubmit}
      // validationSchema={lifecycleIdSchema}
      validateOnBlur={false}
      validateOnChange={false}
      validateOnMount={false}
    >
      {(formikProps: FormikProps<SubmitDatesForm>) => {
        const { errors } = formikProps;
        const flatErrors = flattenErrors(errors);
        return (
          <>
            {Object.keys(errors).length > 0 && (
              <ErrorAlert
                testId="system-intake-errors"
                classNames="margin-top-3"
                heading="Please check and fix the following"
              >
                {Object.keys(flatErrors).map(key => {
                  return (
                    <ErrorAlertMessage
                      key={`Error.${key}`}
                      errorKey={key}
                      message={flatErrors[key]}
                    />
                  );
                })}
              </ErrorAlert>
            )}
            <h1>{t('governanceReviewTeam:dates.heading')}</h1>
            <h2>{t('governanceReviewTeam:dates.subheading')}</h2>
            <div className="tablet:grid-col-9 margin-bottom-7">
              <Form>
                {/* GRT Date Fields */}
                <FieldGroup>
                  <fieldset className="usa-fieldset margin-top-4">
                    <legend className="usa-label margin-bottom-1">
                      {t('governanceReviewTeam:dates.grtDate.label')}
                    </legend>
                    <HelpText className="margin-bottom-1">
                      {t('governanceReviewTeam:dates.grtDate.helpText')}
                    </HelpText>
                    <div
                      className="usa-memorable-date"
                      style={{ marginTop: '-2rem' }}
                    >
                      <div className="usa-form-group usa-form-group--month">
                        <Label htmlFor="Dates-GrtDateMonth">
                          {t('general:date.month')}
                        </Label>
                        <FieldErrorMsg>{flatErrors.grtDateMonth}</FieldErrorMsg>
                        <Field
                          as={TextField}
                          error={!!flatErrors.grtDateMonth}
                          id="Dates-GrtDateMonth"
                          maxLength={2}
                          name="grtDateMonth"
                        />
                      </div>
                      <div className="usa-form-group usa-form-group--day">
                        <Label htmlFor="Dates-GrtDateDay">
                          {t('general:date.day')}
                        </Label>
                        <FieldErrorMsg>{flatErrors.grtDateDay}</FieldErrorMsg>
                        <Field
                          as={TextField}
                          error={!!flatErrors.grtDateDay}
                          id="Dates-GrtDateDay"
                          maxLength={2}
                          name="grtDateDay"
                        />
                      </div>
                      <div className="usa-form-group usa-form-group--year">
                        <Label htmlFor="Dates-GrtDateYear">
                          {t('general:date.year')}
                        </Label>
                        <FieldErrorMsg>{flatErrors.grtDateYear}</FieldErrorMsg>
                        <Field
                          as={TextField}
                          error={!!flatErrors.grtDateYear}
                          id="Dates-GrtDateYear"
                          maxLength={4}
                          name="grtDateYear"
                        />
                      </div>
                    </div>
                  </fieldset>
                </FieldGroup>
                {/* End GRT Date Fields */}
                {/* GRB Date Fields */}
                <FieldGroup>
                  <fieldset className="usa-fieldset margin-top-4">
                    <legend className="usa-label margin-bottom-1">
                      {t('governanceReviewTeam:dates.grbDate.label')}
                    </legend>
                    <HelpText className="margin-bottom-1">
                      {t('governanceReviewTeam:dates.grbDate.helpText')}
                    </HelpText>
                    <div
                      className="usa-memorable-date"
                      style={{ marginTop: '-2rem' }}
                    >
                      <div className="usa-form-group usa-form-group--month">
                        <Label htmlFor="Dates-GrbDateMonth">
                          {t('general:date.month')}
                        </Label>
                        <FieldErrorMsg>{flatErrors.grbDateMonth}</FieldErrorMsg>
                        <Field
                          as={TextField}
                          error={!!flatErrors.grbDateMonth}
                          id="Dates-GrbDateMonth"
                          maxLength={2}
                          name="grbDateMonth"
                        />
                      </div>
                      <div className="usa-form-group usa-form-group--day">
                        <Label htmlFor="Dates-GrbDateDay">
                          {t('general:date.day')}
                        </Label>
                        <FieldErrorMsg>{flatErrors.grbDateDay}</FieldErrorMsg>
                        <Field
                          as={TextField}
                          error={!!flatErrors.grbDateDay}
                          id="Dates-GrbDateDay"
                          maxLength={2}
                          name="grbDateDay"
                        />
                      </div>
                      <div className="usa-form-group usa-form-group--year">
                        <Label htmlFor="Dates-GrbDateYear">
                          {t('general:date.year')}
                        </Label>
                        <FieldErrorMsg>{flatErrors.grbDateYear}</FieldErrorMsg>
                        <Field
                          as={TextField}
                          error={!!flatErrors.grbDateYear}
                          id="Dates-GrbDateYear"
                          maxLength={4}
                          name="grbDateYear"
                        />
                      </div>
                    </div>
                  </fieldset>
                </FieldGroup>
                {/* End GRB Date Fields */}
                <Button className="margin-top-2" type="submit">
                  {t('submit')}
                </Button>
              </Form>
            </div>
          </>
        );
      }}
    </Formik>
  );
};

export default Dates;
