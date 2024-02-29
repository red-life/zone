package resolver

import (
	"fmt"
)

type ResolverService struct {
	repo  ResolverRepository
	cache Cache
}

func (r *ResolverService) ResolveA(domain string) (Answer[ARecord], error) {
	//TODO: less code duplication
	cacheKey := fmt.Sprintf("record:%s:%s", domain, "A")
	value, err := r.cache.Get(cacheKey)
	if err == nil {
		answer := MustFromString[Answer[ARecord]](value)
		return answer, nil
	}
	record, err := r.repo.FindRecord(domain, "A")
	if err != nil {
		return Answer[ARecord]{}, err
	}
	answer := Answer[ARecord]{
		TTL:   record.TTL,
		Value: MustFromString[ARecord](string(record.Value)),
	}
	err = r.cache.Set(cacheKey, MustString(answer))
	if err != nil {
		return Answer[ARecord]{}, err
	}
	return answer, nil
}

func (r *ResolverService) ResolveAAAA(domain string) (Answer[AAAARecord], error) {
	cacheKey := fmt.Sprintf("record:%s:%s", domain, "AAAA")
	value, err := r.cache.Get(cacheKey)
	if err == nil {
		answer := MustFromString[Answer[AAAARecord]](value)
		return answer, nil
	}
	record, err := r.repo.FindRecord(domain, "A")
	if err != nil {
		return Answer[AAAARecord]{}, err
	}
	answer := Answer[AAAARecord]{
		TTL:   record.TTL,
		Value: MustFromString[AAAARecord](string(record.Value)),
	}
	err = r.cache.Set(cacheKey, MustString(answer))
	if err != nil {
		return Answer[AAAARecord]{}, err
	}
	return answer, nil
}

func (r *ResolverService) ResolveMX(domain string) (Answer[MXRecord], error) {
	cacheKey := fmt.Sprintf("record:%s:%s", domain, "MX")
	value, err := r.cache.Get(cacheKey)
	if err == nil {
		answer := MustFromString[Answer[MXRecord]](value)
		return answer, nil
	}
	record, err := r.repo.FindRecord(domain, "A")
	if err != nil {
		return Answer[MXRecord]{}, err
	}
	answer := Answer[MXRecord]{
		TTL:   record.TTL,
		Value: MustFromString[MXRecord](string(record.Value)),
	}
	err = r.cache.Set(cacheKey, MustString(answer))
	if err != nil {
		return Answer[MXRecord]{}, err
	}
	return answer, nil
}

func (r *ResolverService) ResolveNS(domain string) (Answer[NSRecord], error) {
	cacheKey := fmt.Sprintf("record:%s:%s", domain, "NS")
	value, err := r.cache.Get(cacheKey)
	if err == nil {
		answer := MustFromString[Answer[NSRecord]](value)
		return answer, nil
	}
	record, err := r.repo.FindRecord(domain, "A")
	if err != nil {
		return Answer[NSRecord]{}, err
	}
	answer := Answer[NSRecord]{
		TTL:   record.TTL,
		Value: MustFromString[NSRecord](string(record.Value)),
	}
	err = r.cache.Set(cacheKey, MustString(answer))
	if err != nil {
		return Answer[NSRecord]{}, err
	}
	return answer, nil
}

func (r *ResolverService) ResolveTXT(domain string) (Answer[TXTRecord], error) {
	cacheKey := fmt.Sprintf("record:%s:%s", domain, "TXT")
	value, err := r.cache.Get(cacheKey)
	if err == nil {
		answer := MustFromString[Answer[TXTRecord]](value)
		return answer, nil
	}
	record, err := r.repo.FindRecord(domain, "A")
	if err != nil {
		return Answer[TXTRecord]{}, err
	}
	answer := Answer[TXTRecord]{
		TTL:   record.TTL,
		Value: MustFromString[TXTRecord](string(record.Value)),
	}
	err = r.cache.Set(cacheKey, MustString(answer))
	if err != nil {
		return Answer[TXTRecord]{}, err
	}
	return answer, nil
}

func (r *ResolverService) ResolveCNAME(domain string) (Answer[CNAMERecord], error) {
	cacheKey := fmt.Sprintf("record:%s:%s", domain, "CNAME")
	value, err := r.cache.Get(cacheKey)
	if err == nil {
		answer := MustFromString[Answer[CNAMERecord]](value)
		return answer, nil
	}
	record, err := r.repo.FindRecord(domain, "A")
	if err != nil {
		return Answer[CNAMERecord]{}, err
	}
	answer := Answer[CNAMERecord]{
		TTL:   record.TTL,
		Value: MustFromString[CNAMERecord](string(record.Value)),
	}
	err = r.cache.Set(cacheKey, MustString(answer))
	if err != nil {
		return Answer[CNAMERecord]{}, err
	}
	return answer, nil
}

func (r *ResolverService) ResolvePTR(domain string) (Answer[PTRRecord], error) {
	cacheKey := fmt.Sprintf("record:%s:%s", domain, "PTR")
	value, err := r.cache.Get(cacheKey)
	if err == nil {
		answer := MustFromString[Answer[PTRRecord]](value)
		return answer, nil
	}
	record, err := r.repo.FindRecord(domain, "A")
	if err != nil {
		return Answer[PTRRecord]{}, err
	}
	answer := Answer[PTRRecord]{
		TTL:   record.TTL,
		Value: MustFromString[PTRRecord](string(record.Value)),
	}
	err = r.cache.Set(cacheKey, MustString(answer))
	if err != nil {
		return Answer[PTRRecord]{}, err
	}
	return answer, nil
}
