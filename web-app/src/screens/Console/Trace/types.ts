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

export interface CallStats {
  timeToFirstByte: string;
  rx: number;
  tx: number;
  duration: string;
}

export interface TraceMessage {
  client: string;
  time: string;
  ptime: Date;
  statusCode: number;
  api: string;
  query: string;
  host: string;
  callStats: CallStats;
  path: string;
  statusMsg: string;
  key: number;
}
