import { Injectable } from '@nestjs/common';
import { HttpService } from '@nestjs/axios';
import { RpcClient } from 'src/common/rpc.client';
import { Adventure } from 'src/models/adventure';

@Injectable()
export class AdventureDataClient extends RpcClient<Adventure> {
  constructor(httpService: HttpService) {
    super(httpService, 'http://localhost:8000/v1/adventures');
  }
}
