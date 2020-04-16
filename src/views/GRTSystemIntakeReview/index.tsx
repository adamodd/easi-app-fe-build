import React, { useEffect } from 'react';
import { withAuth } from '@okta/okta-react';
import { useDispatch, useSelector } from 'react-redux';
import { RouteComponentProps } from 'react-router-dom';
import { DateTime } from 'luxon';
import {
  DescriptionList,
  DescriptionTerm,
  DescriptionDefinition
} from 'components/shared/DescriptionGroup';
import convertBoolToYesNo from 'utils/convertBoolToYesNo';
import Header from 'components/Header';
import { getSystemIntake } from '../../actions/systemIntakeActions';
import { AppState } from '../../reducers/rootReducer';

export type SystemIDRouterProps = {
  systemId: string;
};

export type GRTSystemIntakeReviewProps = RouteComponentProps<
  SystemIDRouterProps
> & {
  auth: any;
};

export const GRTSystemIntakeReview = ({
  match,
  auth
}: GRTSystemIntakeReviewProps) => {
  const dispatch = useDispatch();

  useEffect(() => {
    const fetchSystemIntake = async (): Promise<void> => {
      dispatch(getSystemIntake(await match.params.systemId));
    };
    fetchSystemIntake();
  }, [auth, dispatch, match.params.systemId]);
  const systemIntake = useSelector(
    (state: AppState) => state.systemIntake.systemIntake
  );
  const fundingDefinition = () => {
    const isFunded = convertBoolToYesNo(!!systemIntake.fundingSource);
    if (systemIntake.fundingSource) {
      return `${isFunded}, ${systemIntake.fundingSource}`;
    }
    return isFunded;
  };
  const issoDefinition = () => {
    const hasIsso = convertBoolToYesNo(!!systemIntake.isso);
    if (systemIntake.isso) {
      return `${hasIsso}, ${systemIntake.isso}`;
    }
    return hasIsso;
  };
  const anyCollaborators = () => {
    if (
      systemIntake.trbCollaborator ||
      systemIntake.oitSecurityCollaborator ||
      systemIntake.eaCollaborator
    ) {
      return true;
    }
    return false;
  };

  return (
    <div className="system-intake__review margin-bottom-7">
      <Header />
      <div className="grid-container">
        <h1 className="font-heading-xl margin-top-4">CMS Sytem Request</h1>
        {!systemIntake && (
          <h2 className="font-heading-xl">
            System intake with ID: {match.params.systemId} was not found
          </h2>
        )}
        {systemIntake && (
          <div>
            <DescriptionList title="System Request">
              <div className="system-intake__review-row">
                <div>
                  <DescriptionTerm term="Submission Date" />
                  <DescriptionDefinition
                    definition={DateTime.local().toLocaleString(
                      DateTime.DATE_MED
                    )}
                  />
                </div>
                <div>
                  <DescriptionTerm term="Request for" />
                  <DescriptionDefinition
                    definition={systemIntake.processStatus}
                  />
                </div>
              </div>
            </DescriptionList>

            <hr className="system-intake__hr" />
            <h2 className="font-heading-xl">Contact Details</h2>

            <DescriptionList title="Contact Details">
              <div className="system-intake__review-row">
                <div>
                  <DescriptionTerm term="Requester" />
                  <DescriptionDefinition definition={systemIntake.requester} />
                </div>
                <div>
                  <DescriptionTerm term="Requester Component" />
                  <DescriptionDefinition definition={systemIntake.component} />
                </div>
              </div>
              <div className="system-intake__review-row">
                <div>
                  <DescriptionTerm term="CMS Business/Product Owner's Name" />
                  <DescriptionDefinition
                    definition={systemIntake.businessOwner}
                  />
                </div>
                <div>
                  <DescriptionTerm term="Business Owner Component" />
                  <DescriptionDefinition
                    definition={systemIntake.businessOwnerComponent}
                  />
                </div>
              </div>
              <div className="system-intake__review-row">
                <div>
                  <DescriptionTerm term="CMS Project/Product Manager or lead" />
                  <DescriptionDefinition
                    definition={systemIntake.productManager}
                  />
                </div>
                <div>
                  <DescriptionTerm term="Product Manager Component" />
                  <DescriptionDefinition
                    definition={systemIntake.productManagerComponent}
                  />
                </div>
              </div>
              <div className="system-intake__review-row">
                <div>
                  <DescriptionTerm term="Does your project have an Information System Security Officer (ISSO)?" />
                  <DescriptionDefinition definition={issoDefinition()} />
                </div>
                <div>
                  <DescriptionTerm term="Currently collaborating with" />
                  {!anyCollaborators && (
                    <DescriptionDefinition definition="N/A" />
                  )}
                  {!!systemIntake.trbCollaborator && (
                    <DescriptionDefinition
                      key="GovernanceTeam-TRB-Collaborator"
                      definition={`Technical Review Board, ${systemIntake.trbCollaborator}`}
                    />
                  )}
                  {!!systemIntake.oitSecurityCollaborator && (
                    <DescriptionDefinition
                      key="GovernanceTeam-OIT-Collaborator"
                      definition={`OIT's Security and Privacy Group, ${systemIntake.oitSecurityCollaborator}`}
                    />
                  )}
                  {!!systemIntake.eaCollaborator && (
                    <DescriptionDefinition
                      key="GovernanceTeam-EA-Collaborator"
                      definition={`Enterprise Architecture, ${systemIntake.eaCollaborator}`}
                    />
                  )}
                </div>
              </div>
            </DescriptionList>

            <hr className="system-intake__hr" />
            <h2 className="font-heading-xl">Request Details</h2>

            <DescriptionList title="Request Details">
              <div className="system-intake__review-row">
                <div>
                  <DescriptionTerm term="Project Name" />
                  <DescriptionDefinition
                    definition={systemIntake.projectName}
                  />
                </div>
                <div>
                  <DescriptionTerm term="Does the project have funding" />
                  <DescriptionDefinition definition={fundingDefinition()} />
                </div>
              </div>
              <div className="margin-bottom-205 line-height-body-3">
                <div>
                  <DescriptionTerm term="What is your business need?" />
                  <DescriptionDefinition
                    definition={systemIntake.businessNeed}
                  />
                </div>
              </div>
              <div className="margin-bottom-205 line-height-body-3">
                <div>
                  <DescriptionTerm term="How are you thinking of solving it?" />
                  <DescriptionDefinition definition={systemIntake.solution} />
                </div>
              </div>
              <div className="system-intake__review-row">
                <div>
                  <DescriptionTerm term="Where are you in the process?" />
                  <DescriptionDefinition
                    definition={systemIntake.processStatus}
                  />
                </div>
                <div>
                  <DescriptionTerm term="Do you currently have a contract in place?" />
                  <DescriptionDefinition
                    definition={convertBoolToYesNo(
                      !!systemIntake.existingContract
                    )}
                  />
                </div>
              </div>
              <div className="system-intake__review-row">
                <div>
                  <DescriptionTerm term="Do you need Enterprise Architecture (EA) support?" />
                  <DescriptionDefinition
                    definition={convertBoolToYesNo(
                      !!systemIntake.needsEaSupport
                    )}
                  />
                </div>
              </div>
            </DescriptionList>
          </div>
        )}
        <hr className="system-intake__hr" />
        <h2 className="font-heading-xl">What to do after reviewing?</h2>
        <p>Email the requester and:</p>
        <ul className="usa-list">
          <li>Ask them to fill in the business case and submit it</li>
          <li>Tell them what to expect after they submit the business case</li>
          <li>And how to get in touch if they have any questions.</li>
        </ul>
      </div>
    </div>
  );
};

export default withAuth(GRTSystemIntakeReview);