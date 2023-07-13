import { Injectable } from '@nestjs/common';
import { HttpService } from '@nestjs/axios';
import { RpcClient } from 'src/common/rpc.client';
import { Partner } from 'src/models/partner';

@Injectable()
export class PartnerDataClient extends RpcClient<Partner> {
  constructor(httpService: HttpService) {
    super(httpService, 'http://localhost:8001/v1/partners');
  }
}
