import {
  STORE_SYSTEM_INTAKE,
  RETRIEVE_SYSTEM_INTAKE,
  UPDATE_SYSTEM_INTAKE
} from 'constants/systemIntake';
import { SystemIntakeForm, SaveSystemIntakeAction } from 'types/systemIntake';

export function saveSystemIntake(
  id: string,
  formData: SystemIntakeForm
): SaveSystemIntakeAction {
  return {
    type: STORE_SYSTEM_INTAKE,
    id,
    formData
  };
}

type GetSystemIntakeAction = {
  type: string;
  intakeID: string;
};

export function getSystemIntake(intakeID: string): GetSystemIntakeAction {
  return {
    type: RETRIEVE_SYSTEM_INTAKE,
    intakeID
  };
}

type PutSystemIntakeAction = {
  type: string;
  systemIntake: SystemIntakeForm;
};

export function putSystemIntake(
  systemIntake: SystemIntakeForm
): PutSystemIntakeAction {
  return {
    type: UPDATE_SYSTEM_INTAKE,
    systemIntake
  };
}
