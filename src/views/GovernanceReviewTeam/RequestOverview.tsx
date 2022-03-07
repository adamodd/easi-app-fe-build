import React, { useEffect } from 'react';
import { useTranslation } from 'react-i18next';
import { useDispatch, useSelector } from 'react-redux';
import { Link, Route, useParams } from 'react-router-dom';
import { useQuery } from '@apollo/client';
import { IconNavigateBefore } from '@trussworks/react-uswds';
import classnames from 'classnames';

import MainContent from 'components/MainContent';
import PageLoading from 'components/PageLoading';
import AddGRTFeedbackKeepDraftBizCase from 'queries/AddGRTFeedbackKeepDraftBizCase';
import AddGRTFeedbackProgressToFinal from 'queries/AddGRTFeedbackProgressToFinal';
import AddGRTFeedbackRequestBizCaseQuery from 'queries/AddGRTFeedbackRequestBizCaseQuery';
import CreateSystemIntakeActionBusinessCaseNeeded from 'queries/CreateSystemIntakeActionBusinessCaseNeededQuery';
import CreateSystemIntakeActionBusinessCaseNeedsChanges from 'queries/CreateSystemIntakeActionBusinessCaseNeedsChangesQuery';
import CreateSystemIntakeActionGuideReceievedClose from 'queries/CreateSystemIntakeActionGuideReceievedCloseQuery';
import CreateSystemIntakeActionNoGovernanceNeeded from 'queries/CreateSystemIntakeActionNoGovernanceNeededQuery';
import CreateSystemIntakeActionNotItRequest from 'queries/CreateSystemIntakeActionNotItRequestQuery';
import CreateSystemIntakeActionNotRespondingClose from 'queries/CreateSystemIntakeActionNotRespondingCloseQuery';
import CreateSystemIntakeActionReadyForGRT from 'queries/CreateSystemIntakeActionReadyForGRTQuery';
import CreateSystemIntakeActionSendEmail from 'queries/CreateSystemIntakeActionSendEmailQuery';
import GetSystemIntakeQuery from 'queries/GetSystemIntakeQuery';
import {
  GetSystemIntake,
  GetSystemIntakeVariables
} from 'queries/types/GetSystemIntake';
import { AppState } from 'reducers/rootReducer';
import { fetchBusinessCase } from 'types/routines';
import ProvideGRTFeedbackToBusinessOwner from 'views/GovernanceReviewTeam/Actions/ProvideGRTFeedbackToBusinessOwner';
import ProvideGRTRecommendationsToGRB from 'views/GovernanceReviewTeam/Actions/ProvideGRTRecommendationsToGRB';
import NotFound from 'views/NotFound';

import ChooseAction from './Actions/ChooseAction';
import ExtendLifecycleId from './Actions/ExtendLifecycleId';
import IssueLifecycleId from './Actions/IssueLifecycleId';
import RejectIntake from './Actions/RejectIntake';
import SubmitAction from './Actions/SubmitAction';
import BusinessCaseReview from './BusinessCaseReview';
import Dates from './Dates';
import Decision from './Decision';
import IntakeReview from './IntakeReview';
import LifecycleID from './LifecycleID';
import Notes from './Notes';
import Summary from './Summary';

import './index.scss';

const RequestOverview = () => {
  const { t } = useTranslation('governanceReviewTeam');
  const { t: actionsT } = useTranslation('action');
  const dispatch = useDispatch();
  const { systemId, activePage } = useParams<{
    systemId: string;
    activePage: string;
  }>();

  const { loading, data, refetch } = useQuery<
    GetSystemIntake,
    GetSystemIntakeVariables
  >(GetSystemIntakeQuery, {
    variables: {
      id: systemId
    }
  });

  const systemIntake = data?.systemIntake;

  const businessCase = useSelector(
    (state: AppState) => state.businessCase.form
  );

  useEffect(() => {
    if (systemIntake?.businessCaseId) {
      dispatch(fetchBusinessCase(systemIntake.businessCaseId));
    }
  }, [dispatch, systemIntake?.businessCaseId]);

  const getNavLinkClasses = (page: string) =>
    classnames('easi-grt__nav-link', {
      'easi-grt__nav-link--active': page === activePage
    });

  if (!loading && !systemIntake) {
    return <NotFound />;
  }

  return (
    <MainContent className="easi-grt" data-testid="grt-request-overview">
      {systemIntake && (
        <Summary
          id={systemIntake.id}
          requester={systemIntake.requester}
          requestName={systemIntake.requestName || ''}
          requestType={systemIntake.requestType}
          status={systemIntake.status}
          adminLead={systemIntake.adminLead}
          submittedAt={systemIntake.submittedAt}
          lcid={systemIntake.lcid}
        />
      )}
      <section className="grid-container grid-row margin-y-5 ">
        <nav className="tablet:grid-col-2 margin-right-2">
          <ul className="easi-grt__nav-list">
            <li className="display-flex flex-align-center">
              <IconNavigateBefore aria-hidden />
              <Link to="/">{t('back.allRequests')}</Link>
            </li>
            <li>
              <Link
                to={`/governance-review-team/${systemId}/intake-request`}
                aria-label={t('aria.openIntake')}
                className={getNavLinkClasses('intake-request')}
              >
                {t('general:intake')}
              </Link>
            </li>
            <li>
              <Link
                to={`/governance-review-team/${systemId}/business-case`}
                aria-label={t('aria.openBusiness')}
                className={getNavLinkClasses('business-case')}
              >
                {t('general:businessCase')}
              </Link>
            </li>
            <li>
              <Link
                to={`/governance-review-team/${systemId}/decision`}
                aria-label={t('aria.openDecision')}
                className={getNavLinkClasses('decision')}
              >
                {t('decision.title')}
              </Link>
            </li>
            <li>
              <Link
                to={`/governance-review-team/${systemId}/lcid`}
                aria-label={t('aria.openLcid')}
                className={getNavLinkClasses('lcid')}
              >
                {t('lifecycleID.title')}
              </Link>
            </li>
          </ul>
          <hr />
          <ul className="easi-grt__nav-list">
            <li>
              <Link
                to={`/governance-review-team/${systemId}/actions`}
                className={getNavLinkClasses('actions')}
                data-testid="grt-nav-actions-link"
              >
                {t('actions')}
              </Link>
            </li>
            <li>
              <Link
                to={`/governance-review-team/${systemId}/notes`}
                className={getNavLinkClasses('notes')}
              >
                {t('notes.heading')}
              </Link>
            </li>
            <li>
              <Link
                to={`/governance-review-team/${systemId}/dates`}
                className={getNavLinkClasses('dates')}
              >
                {t('dates.heading')}
              </Link>
            </li>
          </ul>
        </nav>
        {loading && (
          <div className="margin-x-auto">
            <PageLoading />
          </div>
        )}
        {!loading && !!systemIntake && (
          <section className="tablet:grid-col-9">
            <Route
              path="/governance-review-team/:systemId/intake-request"
              render={() => {
                return <IntakeReview systemIntake={systemIntake} />;
              }}
            />
            <Route
              path="/governance-review-team/:systemId/business-case"
              render={() => (
                <BusinessCaseReview
                  businessCase={businessCase}
                  grtFeedbacks={systemIntake.grtFeedbacks}
                />
              )}
            />
            <Route
              path="/governance-review-team/:systemId/notes"
              render={() => <Notes />}
            />
            <Route
              path="/governance-review-team/:systemId/dates"
              render={() => {
                return <Dates systemIntake={systemIntake} />;
              }}
            />
            <Route
              path="/governance-review-team/:systemId/decision"
              render={() => <Decision systemIntake={systemIntake} />}
            />
            <Route
              path="/governance-review-team/:systemId/lcid"
              render={() => <LifecycleID systemIntake={systemIntake} />}
            />
            <Route
              path="/governance-review-team/:systemId/actions"
              exact
              render={() => (
                <ChooseAction
                  systemIntake={systemIntake}
                  businessCase={businessCase}
                />
              )}
            />
            <Route
              path="/governance-review-team/:systemId/actions/not-it-request"
              render={() => (
                <SubmitAction
                  query={CreateSystemIntakeActionNotItRequest}
                  actionName={actionsT('actions.notItRequest')}
                />
              )}
            />
            <Route
              path="/governance-review-team/:systemId/actions/need-biz-case"
              render={() => (
                <SubmitAction
                  query={CreateSystemIntakeActionBusinessCaseNeeded}
                  actionName={actionsT('actions.needBizCase')}
                />
              )}
            />
            <Route
              path="/governance-review-team/:systemId/actions/provide-feedback-need-biz-case"
              render={() => (
                <ProvideGRTFeedbackToBusinessOwner
                  query={AddGRTFeedbackRequestBizCaseQuery}
                  actionName={actionsT('actions.provideFeedbackNeedBizCase')}
                />
              )}
            />
            <Route
              path="/governance-review-team/:systemId/actions/provide-feedback-keep-draft"
              render={() => (
                <ProvideGRTFeedbackToBusinessOwner
                  query={AddGRTFeedbackKeepDraftBizCase}
                  actionName={actionsT('actions.provideGrtFeedbackKeepDraft')}
                />
              )}
            />
            <Route
              path="/governance-review-team/:systemId/actions/provide-feedback-need-final"
              render={() => (
                <ProvideGRTFeedbackToBusinessOwner
                  query={AddGRTFeedbackProgressToFinal}
                  actionName={actionsT('actions.provideGrtFeedbackNeedFinal')}
                />
              )}
            />
            <Route
              path="/governance-review-team/:systemId/actions/ready-for-grt"
              render={() => (
                <SubmitAction
                  query={CreateSystemIntakeActionReadyForGRT}
                  actionName={actionsT('actions.readyForGrt')}
                />
              )}
            />
            <Route
              path="/governance-review-team/:systemId/actions/ready-for-grb"
              render={() => <ProvideGRTRecommendationsToGRB />}
            />
            <Route
              path="/governance-review-team/:systemId/actions/biz-case-needs-changes"
              render={() => (
                <SubmitAction
                  query={CreateSystemIntakeActionBusinessCaseNeedsChanges}
                  actionName={actionsT('actions.bizCaseNeedsChanges')}
                />
              )}
            />
            <Route
              path="/governance-review-team/:systemId/actions/no-governance"
              render={() => (
                <SubmitAction
                  query={CreateSystemIntakeActionNoGovernanceNeeded}
                  actionName={actionsT('actions.noGovernance')}
                />
              )}
            />
            <Route
              path="/governance-review-team/:systemId/actions/send-email"
              render={() => (
                <SubmitAction
                  query={CreateSystemIntakeActionSendEmail}
                  actionName={actionsT('actions.sendEmail')}
                />
              )}
            />
            <Route
              path="/governance-review-team/:systemId/actions/guide-received-close"
              render={() => (
                <SubmitAction
                  query={CreateSystemIntakeActionGuideReceievedClose}
                  actionName={actionsT('actions.guideReceivedClose')}
                />
              )}
            />
            <Route
              path="/governance-review-team/:systemId/actions/not-responding-close"
              render={() => (
                <SubmitAction
                  query={CreateSystemIntakeActionNotRespondingClose}
                  actionName={actionsT('actions.notRespondingClose')}
                />
              )}
            />
            <Route
              path="/governance-review-team/:systemId/actions/issue-lcid"
              render={() => <IssueLifecycleId />}
            />

            {/* Only display extend LCID action if status is LCID_ISSUED or there has been an lcid issued in the past */}
            {(data?.systemIntake?.status === 'LCID_ISSUED' ||
              data?.systemIntake?.lcid != null) && (
              <Route
                path="/governance-review-team/:systemId/actions/extend-lcid"
                render={() => (
                  <ExtendLifecycleId
                    lcid={data.systemIntake?.lcid || ''}
                    lcidExpiresAt={data.systemIntake?.lcidExpiresAt || ''}
                    lcidScope={data.systemIntake?.lcidScope || ''}
                    lcidNextSteps={data.systemIntake?.decisionNextSteps || ''}
                    lcidCostBaseline={data.systemIntake?.lcidCostBaseline || ''}
                    onSubmit={refetch}
                  />
                )}
              />
            )}
            <Route
              path="/governance-review-team/:systemId/actions/not-approved"
              render={() => <RejectIntake />}
            />
          </section>
        )}
      </section>
    </MainContent>
  );
};

export default RequestOverview;
