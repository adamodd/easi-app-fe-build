import { PUT_SYSTEM_INTAKE, GET_SYSTEM_INTAKE } from 'constants/systemIntake';
import {
  SystemIntakeForm,
  SaveSystemIntakeAction,
  SystemIntakeID
} from 'types/systemIntake';

// eslint-disable-next-line import/prefer-default-export
export function saveSystemIntake(
  id: string,
  formData: SystemIntakeForm
): SaveSystemIntakeAction {
  return {
    type: PUT_SYSTEM_INTAKE,
    id,
    formData
  };
}

type GetSystemIntakeAction = {
  type: string;
  accessToken: string;
  intakeID: SystemIntakeID;
};

// eslint-disable-next-line import/prefer-default-export
export function getSystemIntake(
  accessToken: string,
  intakeID: SystemIntakeID
): GetSystemIntakeAction {
  return {
    type: GET_SYSTEM_INTAKE,
    accessToken,
    intakeID
  };
}
