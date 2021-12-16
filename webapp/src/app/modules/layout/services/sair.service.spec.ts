import { TestBed } from '@angular/core/testing';

import { SairService } from './sair.service';

describe('SairService', () => {
  let service: SairService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(SairService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
