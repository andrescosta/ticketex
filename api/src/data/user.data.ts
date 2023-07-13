import { Injectable } from '@nestjs/common';
import { HttpService } from '@nestjs/axios';
import { User } from '../models/user';
import { RpcClient } from 'src/common/rpc.client';

@Injectable()
export class UserDataClient extends RpcClient<User> {
  constructor(httpService: HttpService) {
    super(httpService, 'http://localhost:8002/v1/users');
  }
}
