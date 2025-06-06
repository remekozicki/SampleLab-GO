package agh.edu.pl.slpbackend.service;

import agh.edu.pl.slpbackend.dto.AddressDto;
import agh.edu.pl.slpbackend.mapper.AddressMapper;
import agh.edu.pl.slpbackend.model.Address;
import agh.edu.pl.slpbackend.repository.AddressRepository;
import lombok.AllArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.stream.Collectors;

@Service
@Slf4j
@AllArgsConstructor
public class AddressService implements AddressMapper {

    private final AddressRepository addressRepository;

    public List<AddressDto> selectAll() {
        log.info("select all address");

        List<Address> addressesList = addressRepository.findAll();
        return addressesList.stream().map(this::toDto).collect(Collectors.toList());
    }
}
