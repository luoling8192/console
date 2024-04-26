// This file is part of FST Console Server
// Copyright (c) 2021 FST, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

import React, { Fragment, useCallback, useEffect } from "react";
import { useSelector } from "react-redux";
import { useLocation, useParams } from "react-router-dom";
import { api } from "api";
import { AppState, useAppDispatch } from "../../../../store";
import { IAM_SCOPES } from "../../../../common/SecureComponent/permissions";
import { decodeURLString, encodeURLString } from "../../../../common/utils";
import {
  resetMessages,
  setIsVersioned,
  setLoadingLocking,
  setLoadingObjectInfo,
  setLoadingVersioning,
  setLoadingVersions,
  setLockingEnabled,
  setObjectDetailsView,
  setRequestInProgress,
  setSelectedObjectView,
  setVersionsModeEnabled,
} from "../../ObjectBrowser/objectBrowserSlice";
import ListObjects from "../ListBuckets/Objects/ListObjects/ListObjects";
import hasPermission from "../../../../common/SecureComponent/accessControl";
import OBHeader from "../../ObjectBrowser/OBHeader";

const BrowserHandler = () => {
  const dispatch = useAppDispatch();
  const params = useParams();
  const location = useLocation();

  const loadingVersioning = useSelector(
    (state: AppState) => state.objectBrowser.loadingVersioning,
  );

  const rewindEnabled = useSelector(
    (state: AppState) => state.objectBrowser.rewind.rewindEnabled,
  );
  const rewindDate = useSelector(
    (state: AppState) => state.objectBrowser.rewind.dateToRewind,
  );
  const showDeleted = useSelector(
    (state: AppState) => state.objectBrowser.showDeleted,
  );
  const requestInProgress = useSelector(
    (state: AppState) => state.objectBrowser.requestInProgress,
  );
  const loadingLocking = useSelector(
    (state: AppState) => state.objectBrowser.loadingLocking,
  );
  const reloadObjectsList = useSelector(
    (state: AppState) => state.objectBrowser.reloadObjectsList,
  );
  const simplePath = useSelector(
    (state: AppState) => state.objectBrowser.simplePath,
  );
  const anonymousMode = useSelector(
    (state: AppState) => state.system.anonymousMode,
  );
  const selectedBucket = useSelector(
    (state: AppState) => state.objectBrowser.selectedBucket,
  );
  const records = useSelector((state: AppState) => state.objectBrowser.records);

  const bucketName = params.bucketName || "";
  const pathSegment = location.pathname.split(`/browser/${bucketName}/`);
  const internalPaths = pathSegment.length === 2 ? pathSegment[1] : "";

  const initWSRequest = useCallback(
    (path: string) => {
      let currDate = new Date();

      let date = currDate.toISOString();

      if (rewindDate !== null && rewindEnabled) {
        date = rewindDate;
      }

      const payloadData = {
        bucketName,
        path,
        rewindMode: rewindEnabled || showDeleted,
        date: date,
      };

      dispatch({ type: "socket/OBRequest", payload: payloadData });
    },
    [bucketName, showDeleted, rewindDate, rewindEnabled, dispatch],
  );

  // Common path load
  const pathLoad = useCallback(
    (forceLoad: boolean = false) => {
      const decodedInternalPaths = decodeURLString(internalPaths);

      // We exit Versions mode in case of path change
      dispatch(setVersionsModeEnabled({ status: false }));

      let searchPath = decodedInternalPaths;

      if (!decodedInternalPaths.endsWith("/") && decodedInternalPaths !== "") {
        searchPath = `${decodedInternalPaths
          .split("/")
          .slice(0, -1)
          .join("/")}/`;
      }

      if (searchPath === "/") {
        searchPath = "";
      }

      // If the path is different of the actual path or reload objects list is requested, then we initialize a new request to load a new record set.
      if (
        searchPath !== simplePath ||
        bucketName !== selectedBucket ||
        forceLoad
      ) {
        dispatch(setRequestInProgress(true));
        initWSRequest(searchPath);
      }
    },
    [
      internalPaths,
      dispatch,
      simplePath,
      selectedBucket,
      bucketName,
      initWSRequest,
    ],
  );

  useEffect(() => {
    return () => {
      dispatch({ type: "socket/OBCancelLast" });
    };
  }, [dispatch]);

  // Object Details handler
  useEffect(() => {
    const decodedIPaths = decodeURLString(internalPaths);

    dispatch(setLoadingVersioning(true));

    if (decodedIPaths.endsWith("/") || decodedIPaths === "") {
      dispatch(setObjectDetailsView(false));
      dispatch(setSelectedObjectView(null));
      dispatch(setLoadingLocking(true));
    } else {
      dispatch(setLoadingObjectInfo(true));
      dispatch(setObjectDetailsView(true));
      dispatch(setLoadingVersions(true));
      dispatch(
        setSelectedObjectView(
          `${decodedIPaths ? `${encodeURLString(decodedIPaths)}` : ``}`,
        ),
      );
    }
  }, [bucketName, internalPaths, rewindDate, rewindEnabled, dispatch]);

  // Navigation Listing Request
  useEffect(() => {
    pathLoad(false);
  }, [pathLoad]);

  // Reload Handler
  useEffect(() => {
    if (reloadObjectsList && records.length === 0 && !requestInProgress) {
      pathLoad(true);
    }
  }, [reloadObjectsList, records, requestInProgress, pathLoad]);

  const displayListObjects =
    hasPermission(bucketName, [
      IAM_SCOPES.S3_LIST_BUCKET,
      IAM_SCOPES.S3_ALL_LIST_BUCKET,
    ]) || anonymousMode;

  useEffect(() => {
    if (loadingVersioning && !anonymousMode) {
      if (displayListObjects) {
        api.buckets
          .getBucketVersioning(bucketName)
          .then((res) => {
            dispatch(setIsVersioned(res.data));
            dispatch(setLoadingVersioning(false));
          })
          .catch((err) => {
            console.error(
              "Error Getting Object Versioning Status: ",
              err.error.detailedMessage,
            );
            dispatch(setLoadingVersioning(false));
          });
      } else {
        dispatch(setLoadingVersioning(false));
        dispatch(resetMessages());
      }
    }
  }, [
    bucketName,
    loadingVersioning,
    dispatch,
    displayListObjects,
    anonymousMode,
  ]);

  useEffect(() => {
    if (loadingLocking) {
      if (displayListObjects) {
        api.buckets
          .getBucketObjectLockingStatus(bucketName)
          .then((res) => {
            dispatch(setLockingEnabled(res.data.object_locking_enabled));
            dispatch(setLoadingLocking(false));
          })
          .catch((err) => {
            console.error(
              "Error Getting Object Locking Status: ",
              err.error.detailedMessage,
            );
            dispatch(setLoadingLocking(false));
          });
      } else {
        dispatch(resetMessages());
        dispatch(setLoadingLocking(false));
      }
    }
  }, [bucketName, loadingLocking, dispatch, displayListObjects]);

  return (
    <Fragment>
      {!anonymousMode && <OBHeader bucketName={bucketName} />}
      <ListObjects />
    </Fragment>
  );
};

export default BrowserHandler;
