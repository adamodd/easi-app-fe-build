import { SystemIntakeForm } from 'types/systemIntake';

export const prepareSystemIntakeForApi = (
  id: string,
  systemIntake: SystemIntakeForm
) => {
  const getGovernanceCollaborator = (name: string) => {
    const selectedTeam = systemIntake.governanceTeams.teams.find(
      team => team.name === name
    );

    return selectedTeam ? selectedTeam.collaborator : '';
  };

  return {
    id,
    requester: systemIntake.requester.name,
    component: systemIntake.requester.component,
    businessOwner: systemIntake.businessOwner.name,
    businessOwnerComponent: systemIntake.businessOwner.component,
    productManager: systemIntake.productManager.name,
    productManagerComponent: systemIntake.productManager.component,
    isso: systemIntake.isso.name,
    trbCollaborator: getGovernanceCollaborator('Technical Review Board'),
    oitSecurityCollaborator: getGovernanceCollaborator(
      "OIT's Security and Privacy Group"
    ),
    eaCollaborator: getGovernanceCollaborator('Enterprise Architecture'),
    projectName: systemIntake.projectName,
    existingFunding: systemIntake.fundingSource.isFunded,
    fundingSource: systemIntake.fundingSource.fundingNumber,
    businessNeed: systemIntake.businessNeed,
    solution: systemIntake.businessSolution,
    processStatus: systemIntake.currentStage,
    eaSupportRequest: systemIntake.needsEaSupport,
    existingContract: systemIntake.hasContract
  };
};

export const prepareSystemIntakeForApp = () => {};
