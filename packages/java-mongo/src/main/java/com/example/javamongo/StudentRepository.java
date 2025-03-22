package com.example.javamongo;

import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.data.repository.query.Param;
import org.springframework.data.rest.core.annotation.RepositoryRestResource;

import java.util.*;

@RepositoryRestResource(collectionResourceRel = "student", path = "student")
public interface StudentRepository extends MongoRepository<Student, String> {
    List<Student> findByName(@Param("name") String name);
}

