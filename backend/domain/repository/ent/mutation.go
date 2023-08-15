// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/szpp-dev-team/hands-on-todo-app/domain/repository/ent/predicate"
	"github.com/szpp-dev-team/hands-on-todo-app/domain/repository/ent/tag"
	"github.com/szpp-dev-team/hands-on-todo-app/domain/repository/ent/task"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeTag  = "Tag"
	TypeTask = "Task"
)

// TagMutation represents an operation that mutates the Tag nodes in the graph.
type TagMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	created_at    *time.Time
	clearedFields map[string]struct{}
	tasks         map[int]struct{}
	removedtasks  map[int]struct{}
	clearedtasks  bool
	done          bool
	oldValue      func(context.Context) (*Tag, error)
	predicates    []predicate.Tag
}

var _ ent.Mutation = (*TagMutation)(nil)

// tagOption allows management of the mutation configuration using functional options.
type tagOption func(*TagMutation)

// newTagMutation creates new mutation for the Tag entity.
func newTagMutation(c config, op Op, opts ...tagOption) *TagMutation {
	m := &TagMutation{
		config:        c,
		op:            op,
		typ:           TypeTag,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withTagID sets the ID field of the mutation.
func withTagID(id int) tagOption {
	return func(m *TagMutation) {
		var (
			err   error
			once  sync.Once
			value *Tag
		)
		m.oldValue = func(ctx context.Context) (*Tag, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Tag.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withTag sets the old Tag of the mutation.
func withTag(node *Tag) tagOption {
	return func(m *TagMutation) {
		m.oldValue = func(context.Context) (*Tag, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m TagMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m TagMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *TagMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *TagMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Tag.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *TagMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *TagMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Tag entity.
// If the Tag object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TagMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *TagMutation) ResetName() {
	m.name = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *TagMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *TagMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Tag entity.
// If the Tag object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TagMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *TagMutation) ResetCreatedAt() {
	m.created_at = nil
}

// AddTaskIDs adds the "tasks" edge to the Task entity by ids.
func (m *TagMutation) AddTaskIDs(ids ...int) {
	if m.tasks == nil {
		m.tasks = make(map[int]struct{})
	}
	for i := range ids {
		m.tasks[ids[i]] = struct{}{}
	}
}

// ClearTasks clears the "tasks" edge to the Task entity.
func (m *TagMutation) ClearTasks() {
	m.clearedtasks = true
}

// TasksCleared reports if the "tasks" edge to the Task entity was cleared.
func (m *TagMutation) TasksCleared() bool {
	return m.clearedtasks
}

// RemoveTaskIDs removes the "tasks" edge to the Task entity by IDs.
func (m *TagMutation) RemoveTaskIDs(ids ...int) {
	if m.removedtasks == nil {
		m.removedtasks = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.tasks, ids[i])
		m.removedtasks[ids[i]] = struct{}{}
	}
}

// RemovedTasks returns the removed IDs of the "tasks" edge to the Task entity.
func (m *TagMutation) RemovedTasksIDs() (ids []int) {
	for id := range m.removedtasks {
		ids = append(ids, id)
	}
	return
}

// TasksIDs returns the "tasks" edge IDs in the mutation.
func (m *TagMutation) TasksIDs() (ids []int) {
	for id := range m.tasks {
		ids = append(ids, id)
	}
	return
}

// ResetTasks resets all changes to the "tasks" edge.
func (m *TagMutation) ResetTasks() {
	m.tasks = nil
	m.clearedtasks = false
	m.removedtasks = nil
}

// Where appends a list predicates to the TagMutation builder.
func (m *TagMutation) Where(ps ...predicate.Tag) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the TagMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *TagMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Tag, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *TagMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *TagMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Tag).
func (m *TagMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *TagMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.name != nil {
		fields = append(fields, tag.FieldName)
	}
	if m.created_at != nil {
		fields = append(fields, tag.FieldCreatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *TagMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case tag.FieldName:
		return m.Name()
	case tag.FieldCreatedAt:
		return m.CreatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *TagMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case tag.FieldName:
		return m.OldName(ctx)
	case tag.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Tag field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TagMutation) SetField(name string, value ent.Value) error {
	switch name {
	case tag.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case tag.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Tag field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *TagMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *TagMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TagMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Tag numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *TagMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *TagMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *TagMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Tag nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *TagMutation) ResetField(name string) error {
	switch name {
	case tag.FieldName:
		m.ResetName()
		return nil
	case tag.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	}
	return fmt.Errorf("unknown Tag field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *TagMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.tasks != nil {
		edges = append(edges, tag.EdgeTasks)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *TagMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case tag.EdgeTasks:
		ids := make([]ent.Value, 0, len(m.tasks))
		for id := range m.tasks {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *TagMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedtasks != nil {
		edges = append(edges, tag.EdgeTasks)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *TagMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case tag.EdgeTasks:
		ids := make([]ent.Value, 0, len(m.removedtasks))
		for id := range m.removedtasks {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *TagMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedtasks {
		edges = append(edges, tag.EdgeTasks)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *TagMutation) EdgeCleared(name string) bool {
	switch name {
	case tag.EdgeTasks:
		return m.clearedtasks
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *TagMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Tag unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *TagMutation) ResetEdge(name string) error {
	switch name {
	case tag.EdgeTasks:
		m.ResetTasks()
		return nil
	}
	return fmt.Errorf("unknown Tag edge %s", name)
}

// TaskMutation represents an operation that mutates the Task nodes in the graph.
type TaskMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	description   *string
	deadline      *time.Time
	completd_at   *time.Time
	created_at    *time.Time
	updated_at    *time.Time
	clearedFields map[string]struct{}
	tags          map[int]struct{}
	removedtags   map[int]struct{}
	clearedtags   bool
	done          bool
	oldValue      func(context.Context) (*Task, error)
	predicates    []predicate.Task
}

var _ ent.Mutation = (*TaskMutation)(nil)

// taskOption allows management of the mutation configuration using functional options.
type taskOption func(*TaskMutation)

// newTaskMutation creates new mutation for the Task entity.
func newTaskMutation(c config, op Op, opts ...taskOption) *TaskMutation {
	m := &TaskMutation{
		config:        c,
		op:            op,
		typ:           TypeTask,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withTaskID sets the ID field of the mutation.
func withTaskID(id int) taskOption {
	return func(m *TaskMutation) {
		var (
			err   error
			once  sync.Once
			value *Task
		)
		m.oldValue = func(ctx context.Context) (*Task, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Task.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withTask sets the old Task of the mutation.
func withTask(node *Task) taskOption {
	return func(m *TaskMutation) {
		m.oldValue = func(context.Context) (*Task, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m TaskMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m TaskMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Task entities.
func (m *TaskMutation) SetID(id int) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *TaskMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *TaskMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Task.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *TaskMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *TaskMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Task entity.
// If the Task object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TaskMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *TaskMutation) ResetName() {
	m.name = nil
}

// SetDescription sets the "description" field.
func (m *TaskMutation) SetDescription(s string) {
	m.description = &s
}

// Description returns the value of the "description" field in the mutation.
func (m *TaskMutation) Description() (r string, exists bool) {
	v := m.description
	if v == nil {
		return
	}
	return *v, true
}

// OldDescription returns the old "description" field's value of the Task entity.
// If the Task object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TaskMutation) OldDescription(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDescription is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDescription requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDescription: %w", err)
	}
	return oldValue.Description, nil
}

// ClearDescription clears the value of the "description" field.
func (m *TaskMutation) ClearDescription() {
	m.description = nil
	m.clearedFields[task.FieldDescription] = struct{}{}
}

// DescriptionCleared returns if the "description" field was cleared in this mutation.
func (m *TaskMutation) DescriptionCleared() bool {
	_, ok := m.clearedFields[task.FieldDescription]
	return ok
}

// ResetDescription resets all changes to the "description" field.
func (m *TaskMutation) ResetDescription() {
	m.description = nil
	delete(m.clearedFields, task.FieldDescription)
}

// SetDeadline sets the "deadline" field.
func (m *TaskMutation) SetDeadline(t time.Time) {
	m.deadline = &t
}

// Deadline returns the value of the "deadline" field in the mutation.
func (m *TaskMutation) Deadline() (r time.Time, exists bool) {
	v := m.deadline
	if v == nil {
		return
	}
	return *v, true
}

// OldDeadline returns the old "deadline" field's value of the Task entity.
// If the Task object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TaskMutation) OldDeadline(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDeadline is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDeadline requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDeadline: %w", err)
	}
	return oldValue.Deadline, nil
}

// ClearDeadline clears the value of the "deadline" field.
func (m *TaskMutation) ClearDeadline() {
	m.deadline = nil
	m.clearedFields[task.FieldDeadline] = struct{}{}
}

// DeadlineCleared returns if the "deadline" field was cleared in this mutation.
func (m *TaskMutation) DeadlineCleared() bool {
	_, ok := m.clearedFields[task.FieldDeadline]
	return ok
}

// ResetDeadline resets all changes to the "deadline" field.
func (m *TaskMutation) ResetDeadline() {
	m.deadline = nil
	delete(m.clearedFields, task.FieldDeadline)
}

// SetCompletdAt sets the "completd_at" field.
func (m *TaskMutation) SetCompletdAt(t time.Time) {
	m.completd_at = &t
}

// CompletdAt returns the value of the "completd_at" field in the mutation.
func (m *TaskMutation) CompletdAt() (r time.Time, exists bool) {
	v := m.completd_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCompletdAt returns the old "completd_at" field's value of the Task entity.
// If the Task object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TaskMutation) OldCompletdAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCompletdAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCompletdAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCompletdAt: %w", err)
	}
	return oldValue.CompletdAt, nil
}

// ClearCompletdAt clears the value of the "completd_at" field.
func (m *TaskMutation) ClearCompletdAt() {
	m.completd_at = nil
	m.clearedFields[task.FieldCompletdAt] = struct{}{}
}

// CompletdAtCleared returns if the "completd_at" field was cleared in this mutation.
func (m *TaskMutation) CompletdAtCleared() bool {
	_, ok := m.clearedFields[task.FieldCompletdAt]
	return ok
}

// ResetCompletdAt resets all changes to the "completd_at" field.
func (m *TaskMutation) ResetCompletdAt() {
	m.completd_at = nil
	delete(m.clearedFields, task.FieldCompletdAt)
}

// SetCreatedAt sets the "created_at" field.
func (m *TaskMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *TaskMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Task entity.
// If the Task object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TaskMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *TaskMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *TaskMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *TaskMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Task entity.
// If the Task object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TaskMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (m *TaskMutation) ClearUpdatedAt() {
	m.updated_at = nil
	m.clearedFields[task.FieldUpdatedAt] = struct{}{}
}

// UpdatedAtCleared returns if the "updated_at" field was cleared in this mutation.
func (m *TaskMutation) UpdatedAtCleared() bool {
	_, ok := m.clearedFields[task.FieldUpdatedAt]
	return ok
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *TaskMutation) ResetUpdatedAt() {
	m.updated_at = nil
	delete(m.clearedFields, task.FieldUpdatedAt)
}

// AddTagIDs adds the "tags" edge to the Tag entity by ids.
func (m *TaskMutation) AddTagIDs(ids ...int) {
	if m.tags == nil {
		m.tags = make(map[int]struct{})
	}
	for i := range ids {
		m.tags[ids[i]] = struct{}{}
	}
}

// ClearTags clears the "tags" edge to the Tag entity.
func (m *TaskMutation) ClearTags() {
	m.clearedtags = true
}

// TagsCleared reports if the "tags" edge to the Tag entity was cleared.
func (m *TaskMutation) TagsCleared() bool {
	return m.clearedtags
}

// RemoveTagIDs removes the "tags" edge to the Tag entity by IDs.
func (m *TaskMutation) RemoveTagIDs(ids ...int) {
	if m.removedtags == nil {
		m.removedtags = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.tags, ids[i])
		m.removedtags[ids[i]] = struct{}{}
	}
}

// RemovedTags returns the removed IDs of the "tags" edge to the Tag entity.
func (m *TaskMutation) RemovedTagsIDs() (ids []int) {
	for id := range m.removedtags {
		ids = append(ids, id)
	}
	return
}

// TagsIDs returns the "tags" edge IDs in the mutation.
func (m *TaskMutation) TagsIDs() (ids []int) {
	for id := range m.tags {
		ids = append(ids, id)
	}
	return
}

// ResetTags resets all changes to the "tags" edge.
func (m *TaskMutation) ResetTags() {
	m.tags = nil
	m.clearedtags = false
	m.removedtags = nil
}

// Where appends a list predicates to the TaskMutation builder.
func (m *TaskMutation) Where(ps ...predicate.Task) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the TaskMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *TaskMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Task, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *TaskMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *TaskMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Task).
func (m *TaskMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *TaskMutation) Fields() []string {
	fields := make([]string, 0, 6)
	if m.name != nil {
		fields = append(fields, task.FieldName)
	}
	if m.description != nil {
		fields = append(fields, task.FieldDescription)
	}
	if m.deadline != nil {
		fields = append(fields, task.FieldDeadline)
	}
	if m.completd_at != nil {
		fields = append(fields, task.FieldCompletdAt)
	}
	if m.created_at != nil {
		fields = append(fields, task.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, task.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *TaskMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case task.FieldName:
		return m.Name()
	case task.FieldDescription:
		return m.Description()
	case task.FieldDeadline:
		return m.Deadline()
	case task.FieldCompletdAt:
		return m.CompletdAt()
	case task.FieldCreatedAt:
		return m.CreatedAt()
	case task.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *TaskMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case task.FieldName:
		return m.OldName(ctx)
	case task.FieldDescription:
		return m.OldDescription(ctx)
	case task.FieldDeadline:
		return m.OldDeadline(ctx)
	case task.FieldCompletdAt:
		return m.OldCompletdAt(ctx)
	case task.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case task.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Task field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TaskMutation) SetField(name string, value ent.Value) error {
	switch name {
	case task.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case task.FieldDescription:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDescription(v)
		return nil
	case task.FieldDeadline:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDeadline(v)
		return nil
	case task.FieldCompletdAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCompletdAt(v)
		return nil
	case task.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case task.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Task field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *TaskMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *TaskMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TaskMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Task numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *TaskMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(task.FieldDescription) {
		fields = append(fields, task.FieldDescription)
	}
	if m.FieldCleared(task.FieldDeadline) {
		fields = append(fields, task.FieldDeadline)
	}
	if m.FieldCleared(task.FieldCompletdAt) {
		fields = append(fields, task.FieldCompletdAt)
	}
	if m.FieldCleared(task.FieldUpdatedAt) {
		fields = append(fields, task.FieldUpdatedAt)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *TaskMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *TaskMutation) ClearField(name string) error {
	switch name {
	case task.FieldDescription:
		m.ClearDescription()
		return nil
	case task.FieldDeadline:
		m.ClearDeadline()
		return nil
	case task.FieldCompletdAt:
		m.ClearCompletdAt()
		return nil
	case task.FieldUpdatedAt:
		m.ClearUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown Task nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *TaskMutation) ResetField(name string) error {
	switch name {
	case task.FieldName:
		m.ResetName()
		return nil
	case task.FieldDescription:
		m.ResetDescription()
		return nil
	case task.FieldDeadline:
		m.ResetDeadline()
		return nil
	case task.FieldCompletdAt:
		m.ResetCompletdAt()
		return nil
	case task.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case task.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown Task field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *TaskMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.tags != nil {
		edges = append(edges, task.EdgeTags)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *TaskMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case task.EdgeTags:
		ids := make([]ent.Value, 0, len(m.tags))
		for id := range m.tags {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *TaskMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedtags != nil {
		edges = append(edges, task.EdgeTags)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *TaskMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case task.EdgeTags:
		ids := make([]ent.Value, 0, len(m.removedtags))
		for id := range m.removedtags {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *TaskMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedtags {
		edges = append(edges, task.EdgeTags)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *TaskMutation) EdgeCleared(name string) bool {
	switch name {
	case task.EdgeTags:
		return m.clearedtags
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *TaskMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Task unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *TaskMutation) ResetEdge(name string) error {
	switch name {
	case task.EdgeTags:
		m.ResetTags()
		return nil
	}
	return fmt.Errorf("unknown Task edge %s", name)
}
