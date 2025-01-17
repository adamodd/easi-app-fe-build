import React, { Fragment } from 'react';
import { useTranslation } from 'react-i18next';
import { Field, FieldArray, FormikProps, getIn } from 'formik';

import CheckboxField from 'components/shared/CheckboxField';
import FieldErrorMsg from 'components/shared/FieldErrorMsg';
import FieldGroup from 'components/shared/FieldGroup';
import Label from 'components/shared/Label';
import TextField from 'components/shared/TextField';
import cmsGovernanceTeams from 'constants/enums/cmsGovernanceTeams';

import { ContactDetailsForm } from './index';

type GovernanceTeamOptionsProps = {
  formikProps: FormikProps<ContactDetailsForm>;
};

const GovernanceTeamOptions = ({ formikProps }: GovernanceTeamOptionsProps) => {
  const { values, errors } = formikProps;
  const teams = values.governanceTeams.teams || [];
  const { t } = useTranslation('intake');

  return (
    <FieldArray name="governanceTeams.teams">
      {arrayHelpers => (
        <fieldset className="margin-0 padding-0 border-0">
          <legend className="usa-sr-only">
            {t('contactDetails.governanceTeam.helpText')}
          </legend>
          {cmsGovernanceTeams.map((team: any, index: number) => {
            const isChecked = teams
              .map(govTeam => govTeam.name)
              .includes(team.value);
            return (
              <Fragment key={team.key}>
                <CheckboxField
                  checked={isChecked}
                  disabled={values.governanceTeams.isPresent !== true}
                  id={`governanceTeam-${team.key}`}
                  label={team.label}
                  name={`governanceTeams.teams.${index}`}
                  onBlur={() => {}}
                  onChange={e => {
                    if (e.target.checked) {
                      arrayHelpers.push({
                        name: e.target.value,
                        key: team.key,
                        collaborator: ''
                      });
                    } else {
                      const removeIndex = values.governanceTeams.teams
                        ?.map(govTeam => govTeam.name)
                        .indexOf(e.target.value);
                      if (typeof removeIndex === 'number' && removeIndex > -1) {
                        arrayHelpers.remove(removeIndex);
                      }
                    }
                  }}
                  value={team.value}
                  inputProps={{
                    'aria-expanded': isChecked,
                    'aria-controls': `${team.key}-Collaborator`
                  }}
                />
                {teams.map((govTeam, idx) => {
                  const { key } = team;
                  if (team.value === govTeam.name) {
                    return (
                      <div
                        key={`${key}-Collaborator`}
                        id={`${key}-Collaborator`}
                        className="width-card-lg margin-top-neg-2 margin-left-4 margin-bottom-2"
                      >
                        <FieldGroup
                          scrollElement={`governanceTeams.teams.${idx}.collaborator`}
                          error={false}
                        >
                          <Label
                            htmlFor={`IntakeForm-${key}-Collaborator`}
                            aria-label={t(
                              `Enter ${team.acronym} Collaborator Name`
                            )}
                          >
                            {t(`${team.acronym} Collaborator Name`)}
                          </Label>
                          {errors.governanceTeams &&
                            errors.governanceTeams.teams &&
                            errors.governanceTeams.teams[idx] && (
                              <FieldErrorMsg>
                                {getIn(
                                  errors,
                                  `governanceTeams.teams.${idx}.collaborator`
                                )}
                              </FieldErrorMsg>
                            )}
                          <Field
                            as={TextField}
                            error={getIn(
                              errors,
                              `governanceTeams.teams.${idx}.collaborator`
                            )}
                            id={`IntakeForm-${key}-Collaborator`}
                            maxLength={50}
                            name={`governanceTeams.teams.${idx}.collaborator`}
                          />
                        </FieldGroup>
                      </div>
                    );
                  }
                  return null;
                })}
              </Fragment>
            );
          })}
        </fieldset>
      )}
    </FieldArray>
  );
};

export default GovernanceTeamOptions;
