interface Capacity {
  type: string;
  availability: number;
}

export interface ReservationMetadata {
  adventure_id: string;
  status: number;
  capacities: Array<Capacity>;
}
