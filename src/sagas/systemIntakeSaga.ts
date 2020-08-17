import axios from 'axios';
import { Action } from 'redux-actions';
import { call, put, takeLatest } from 'redux-saga/effects';

import { prepareSystemIntakeForApi } from 'data/systemIntake';
import { updateLastActiveAt } from 'reducers/authReducer';
import {
  archiveSystemIntake,
  fetchSystemIntake,
  postSystemIntake,
  reviewSystemIntake,
  saveSystemIntake,
  submitSystemIntake
} from 'types/routines';
import { SystemIntakeForm } from 'types/systemIntake';

function putSystemIntakeRequest(formData: SystemIntakeForm) {
  // Make API save request
  const data = prepareSystemIntakeForApi(formData);
  return axios.put(`${process.env.REACT_APP_API_ADDRESS}/system_intake`, data);
}

function postSystemIntakeRequest(formData: SystemIntakeForm) {
  const data = prepareSystemIntakeForApi(formData);
  return axios.post(`${process.env.REACT_APP_API_ADDRESS}/system_intake`, data);
}

function* createSystemIntake(action: Action<any>) {
  try {
    yield put(postSystemIntake.request());
    const response = yield call(postSystemIntakeRequest, action.payload);
    yield put(postSystemIntake.success(response.data));
  } catch (error) {
    yield put(postSystemIntake.failure(error.message));
  } finally {
    yield put(postSystemIntake.fulfill());
    yield put(updateLastActiveAt);
  }
}

function* putSystemIntake(action: Action<any>) {
  try {
    yield put(saveSystemIntake.request());
    const response = yield call(putSystemIntakeRequest, action.payload);
    yield put(saveSystemIntake.success(response.data));
  } catch (error) {
    yield put(saveSystemIntake.failure(error.message));
  } finally {
    yield put(saveSystemIntake.fulfill());
    yield put(updateLastActiveAt);
  }
}

function getSystemIntakeRequest(id: string) {
  return axios.get(`${process.env.REACT_APP_API_ADDRESS}/system_intake/${id}`);
}

function* getSystemIntake(action: Action<any>) {
  try {
    yield put(fetchSystemIntake.request());
    const response = yield call(getSystemIntakeRequest, action.payload);
    yield put(fetchSystemIntake.success(response.data));
  } catch (error) {
    yield put(fetchSystemIntake.failure(error.message));
  } finally {
    yield put(fetchSystemIntake.fulfill());
    yield put(updateLastActiveAt);
  }
}

function* completeSystemIntake(action: Action<any>) {
  try {
    yield put(submitSystemIntake.request());
    const response = yield call(putSystemIntakeRequest, {
      ...action.payload,
      status: 'SUBMITTED'
    });
    yield put(submitSystemIntake.success(response.data));
  } catch (error) {
    yield put(submitSystemIntake.failure(error.message));
  } finally {
    yield put(submitSystemIntake.fulfill());
    yield put(updateLastActiveAt);
  }
}

function* submitSystemIntakeReview(action: Action<any>) {
  try {
    yield put(reviewSystemIntake.request());
    const response = yield call(putSystemIntakeRequest, action.payload);
    yield put(reviewSystemIntake.success(response.data));
  } catch (error) {
    yield put(reviewSystemIntake.failure(error.message));
  } finally {
    yield put(reviewSystemIntake.fulfill());
  }
}

function deleteSystemIntakeRequest(id: string) {
  return axios.delete(
    `${process.env.REACT_APP_API_ADDRESS}/system_intake/${id}`
  );
}

function* deleteSystemIntake(action: Action<any>) {
  try {
    yield put(archiveSystemIntake.request());
    const response = yield call(
      deleteSystemIntakeRequest,
      action.payload.intakeId
    );
    yield put(archiveSystemIntake.success(response.data));
    action.payload.redirect();
  } catch (error) {
    yield put(archiveSystemIntake.failure(error.message));
  } finally {
    yield put(archiveSystemIntake.fulfill());
    yield put(updateLastActiveAt);
  }
}

export default function* systemIntakeSaga() {
  yield takeLatest(fetchSystemIntake.TRIGGER, getSystemIntake);
  yield takeLatest(saveSystemIntake.TRIGGER, putSystemIntake);
  yield takeLatest(submitSystemIntake.TRIGGER, completeSystemIntake);
  yield takeLatest(postSystemIntake.TRIGGER, createSystemIntake);
  yield takeLatest(reviewSystemIntake.TRIGGER, submitSystemIntakeReview);
  yield takeLatest(archiveSystemIntake.TRIGGER, deleteSystemIntake);
}
